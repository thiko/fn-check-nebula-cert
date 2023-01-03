package common

import "time"

type ObjectStoreRepository interface {
	ListFiles(bucketName string) ([]string, error)
	Download(bucketName string, objectName string) ([]byte, error)
	Upload(bucketName string, content []byte) error
}

// internal config struct
type RunConfig struct {
	Repository            ObjectStoreRepository
	Mode                  string
	CertificateBucketName string
	ResultBucketName      string
}

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

type Worker interface {
	Work(runconfig *RunConfig)
}
