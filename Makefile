CERT_PATH ?=``

fn-nebula-cert-check-dev:
	go run main.go $(CERT_PATH)

fn-nebula-cert-check-test:
	go test -v ./...

fn-nebula-cert-check-test-coverage:
	go test -cover  ./...