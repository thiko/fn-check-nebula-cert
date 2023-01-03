build-aws:
	go build ./cmd/aws

build-lambda:
	GOOS=linux GOARCH=amd64 go build -o aws_lambda ./cmd/aws
	chmod +x aws_lambda
	zip main.zip aws_lambda

exec-local:
	go run . -cert-bucket=my-test-bucket -mode=local

exec-test:
	go test -v ./...

exec-test-coverage:
	go test -cover  ./...