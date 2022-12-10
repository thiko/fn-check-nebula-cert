package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"time"

	"github.com/slackhq/nebula/cert"
)

type ValidationResults struct {
	ValidationResults []ValidationResult `json:"results"`
	ExecutedAt        time.Time          `json:"executed_at"`
}
type ValidationResult struct {
	Name       string    `json:"name"`
	Signature  []byte    `json:"signature"`
	ValidFrom  time.Time `json:"valid_from"`
	ValidUntil time.Time `json:"valid_until"`
	Expired    bool      `json:"is_expired"`
}

type ObjectStoreRepository interface {
	ListFiles(bucketName string) ([]string, error)
	Download(bucketName string, objectName string) ([]byte, error)
	Upload(bucketName string, content []byte) error
}

// internal config struct
type runConfig struct {
	repository            ObjectStoreRepository
	mode                  string
	certificateBucketName string
	resultBucketName      string
}

func main() {

	modePtr := flag.String("mode", "local", "Execution mode: local, aws or gcp")

	certBucketPtr := flag.String("cert-bucket", "nebula-cert-bucket", "Defines the certificate bucket name")
	resultBucketPtr := flag.String("result-bucket", "nebula-result-bucket", "Defines the result bucket name")

	flag.Parse()

	config, err := runConfigurationFromArguments(modePtr, certBucketPtr, resultBucketPtr)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Executing function in mode: `%s`", *modePtr)

	objects, err := config.repository.ListFiles(config.certificateBucketName)
	if err != nil {
		log.Fatalf("Unable to list files from bucket `%s`. Error: %s", config.certificateBucketName, err)
	}

	log.Printf("Found files in bucket (`%s`): %v", config.certificateBucketName, objects)
	validationResultList := make([]ValidationResult, 0)

	for i := 0; i < len(objects); i++ {
		var currObj = objects[i]
		log.Printf("Start download of: `%s`", currObj)

		certFile, err := config.repository.Download(config.certificateBucketName, currObj)
		if err != nil {
			log.Printf("WARN: Skip file; Unable to download file: `%s/%s`; %s", config.certificateBucketName, currObj, err)
			continue
		}

		nebulaCertificate, err := ValidateCertificate(certFile)
		if err != nil {
			log.Printf("WARN: Skip file; Unable to parse certificate: `%s/%s`; %s", config.certificateBucketName, currObj, err)
			continue
		}

		validationResultList = append(validationResultList, *nebulaCertificate)
	}

	totalResult := ValidationResults{
		ValidationResults: validationResultList,
		ExecutedAt:        time.Now(),
	}

	err = storeResultOnStorage(config.repository, config.resultBucketName, &totalResult)
	if err != nil {
		log.Fatalf("Error while uploading results: %s", err)
	}
	log.Print("Result successfuly stored")
}

func ValidateCertificate(fileContent []byte) (*ValidationResult, error) {
	nebulaCertificate, _, err := cert.UnmarshalNebulaCertificateFromPEM(fileContent)

	if err != nil {
		log.Printf("Error: Unable to parse certificate:`%s`", err)
		return nil, err
	}

	return &ValidationResult{
		Name:       nebulaCertificate.Details.Name,
		Signature:  nebulaCertificate.Signature,
		ValidFrom:  nebulaCertificate.Details.NotBefore,
		ValidUntil: nebulaCertificate.Details.NotAfter,
		Expired:    nebulaCertificate.Expired(time.Now()),
	}, nil
}

func storeResultOnStorage(repository ObjectStoreRepository, resultBucketName string, validationResults *ValidationResults) error {

	jsonCert, err := json.Marshal(validationResults)
	if err != nil {
		log.Fatalf("Unable create json from validation result: %v; Error: `%s`", validationResults, err)
	}

	return repository.Upload(resultBucketName, jsonCert)
}

func runConfigurationFromArguments(modePtr *string, certBucketPtr *string, resultBucketPtr *string) (*runConfig, error) {

	var repository ObjectStoreRepository

	switch *modePtr {
	case "local":
		{
			repository = &LocalRepository{}
		}
	case "aws":
		{
			repository = &AwsRepository{}
		}
	case "gcp":
		{
			repository = &GcpRepository{}
		}
	default:
		return nil, fmt.Errorf("unknown execution mode `%s`", *modePtr)
	}

	return &runConfig{
		repository:            repository,
		mode:                  *modePtr,
		certificateBucketName: *certBucketPtr,
		resultBucketName:      *certBucketPtr,
	}, nil
}
