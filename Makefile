build-aws:
	go build -o bin/aws_lambda_manual ./cmd/aws 

build-lambda:
	GOOS=linux GOARCH=amd64 go build -o bin/aws_lambda ./cmd/aws
	chmod +x bin/aws_lambda
	zip bin/aws_lambda.zip bin/aws_lambda
	rm -f bin/aws_lambda

exec-local:
	go run . -cert_bucket=my-test-bucket -mode=local

exec-test:
	go test -v ./...

exec-test-coverage:
	go test -cover  ./...