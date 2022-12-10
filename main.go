package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
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

func main() {

	if len(os.Args[1:]) == 0 {
		log.Fatalf("Mandatory field missing. You have to pass the path to the certificate file.")
	}

	certFilePath := os.Args[1]
	fileContent, err := ioutil.ReadFile(certFilePath)
	if err != nil {
		log.Fatalf("Unable to read file: `%s`; %s", certFilePath, err)
	}

	nebulaCertificate, err := ValidateCertificate(fileContent)

	if err != nil {
		log.Fatalf("Unable to parse certificate: %s; Error: `%s`", certFilePath, err)
	}

	results := ValidationResults{

		ValidationResults: []ValidationResult{*nebulaCertificate},
		ExecutedAt:        time.Now(),
	}

	storeResultOnStorage(&results)
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

func storeResultOnStorage(validationResults *ValidationResults) {

	jsonCert, err := json.Marshal(validationResults)
	if err != nil {
		log.Fatalf("Unable create json from validation result: %v; Error: `%s`", validationResults, err)
	}

	log.Print(string(jsonCert))
}
