package main

import (
	"fmt"
	"log"
)

const DUMMY_BUCKET = "lulu"

const validFileName = "testfile_valid.crt"
const expiredFileName = "testfile_expired.crt"

// expires at 2027-11-11 as well... stupid test :)
const validFileContent = `-----BEGIN NEBULA CERTIFICATE-----
CkAKDm5lYnVsYSByb290IGNhKJfap9AFMJfg1+YGOiCUQGByMuNRhIlQBOyzXWbL
vcKBwDhov900phEfJ5DN3kABEkDCq5R8qBiu8sl54yVfgRcQXEDt3cHr8UTSLszv
bzBEr00kERQxxTzTsH8cpYEgRoipvmExvg8WP8NdAJEYJosB
-----END NEBULA CERTIFICATE-----`
const expiredFileContent = `
# expired certificate
-----BEGIN NEBULA CERTIFICATE-----
CjkKB2V4cGlyZWQouPmWjQYwufmWjQY6ILCRaoCkJlqHgv5jfDN4lzLHBvDzaQm4
vZxfu144hmgjQAESQG4qlnZi8DncvD/LDZnLgJHOaX1DWCHHEh59epVsC+BNgTie
WH1M9n4O7cFtGlM6sJJOS+rCVVEJ3ABS7+MPdQs=
-----END NEBULA CERTIFICATE-----
`

type LocalRepository struct{}

func (LocalRepository) ListFiles(bucketName string) ([]string, error) {
	return []string{validFileName, expiredFileName}, nil
}

func (*LocalRepository) Download(bucketName string, objectName string) ([]byte, error) {
	if objectName == validFileName {
		return []byte(validFileContent), nil
	} else if objectName == expiredFileName {
		return []byte(expiredFileContent), nil
	}

	return []byte{}, fmt.Errorf("unknown object name `%s`", objectName)
}

func (*LocalRepository) Upload(bucketName string, content []byte) error {
	log.Printf("Fake upload to bucket `%s`", bucketName)
	log.Printf("Result: `%s`", string(content))

	return nil
}
