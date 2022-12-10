
exec-local:
	go run . -cert-bucket=my-test-bucket -mode=local

exec-test:
	go test -v ./...

exec-test-coverage:
	go test -cover  ./...