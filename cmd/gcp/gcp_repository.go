package gcp

type GcpRepository struct{}

func (*GcpRepository) ListFiles(bucketName string) ([]string, error) {
	return []string{}, nil
}

func (*GcpRepository) Download(bucketName string, objectName string) ([]byte, error) {
	return []byte{}, nil
}

func (*GcpRepository) Upload(bucketName string, content []byte) error {
	return nil
}
