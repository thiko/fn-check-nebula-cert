package main

type AwsRepository struct{}

func (*AwsRepository) ListFiles(bucketName string) ([]string, error) {
	return []string{}, nil
}

func (*AwsRepository) Download(bucketName string, objectName string) ([]byte, error) {
	return []byte{}, nil
}

func (*AwsRepository) Upload(bucketName string, content []byte) error {
	return nil
}
