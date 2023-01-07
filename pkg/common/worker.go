package common

import (
	"encoding/json"
	"log"
	"time"

	"github.com/slackhq/nebula/cert"
)

type DefaultWorker struct{}

func (*DefaultWorker) Work(config *RunConfig) {

	log.Printf("Executing function in mode: `%s`. Using certificate bucket: `%s` and store results in: `%s`", config.Mode, config.CertificateBucketName, config.ResultBucketName)

	objects, err := config.Repository.ListFiles(config.CertificateBucketName)
	if err != nil {
		log.Fatalf("Unable to list files from bucket `%s`. Error: %s", config.CertificateBucketName, err)
	}

	log.Printf("Found files in bucket (`%s`): %v", config.CertificateBucketName, objects)
	validationResultList := make([]ValidationResult, 0)

	for i := 0; i < len(objects); i++ {
		var currObj = objects[i]
		log.Printf("Start download of: `%s`", currObj)

		certFile, err := config.Repository.Download(config.CertificateBucketName, currObj)
		if err != nil {
			log.Printf("WARN: Skip file; Unable to download file: `%s/%s`; %s", config.CertificateBucketName, currObj, err)
			continue
		}

		nebulaCertificate, err := ValidateCertificate(certFile)
		if err != nil {
			log.Printf("WARN: Skip file; Unable to parse certificate: `%s/%s`; %s", config.CertificateBucketName, currObj, err)
			continue
		}

		validationResultList = append(validationResultList, *nebulaCertificate)
	}

	totalResult := ValidationResults{
		ValidationResults: validationResultList,
		ExecutedAt:        time.Now(),
	}

	err = storeResultOnStorage(config.Repository, config.ResultBucketName, &totalResult)
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
