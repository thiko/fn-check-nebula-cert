# Check Nebula Cert Function
- Reads all .crt files from a defined Bucket
- Validates each .crt file including expiration check
- Stores the result in a file and stores that file in the same bucket. Its using the same file on each invokation (overwrite it).


## Execution

Given Makefile supports different execution & testing modes.

`make fn-nebula-cert-check-dev CERT_PATH="path/to/cert"`: Runs the application. Accepts the path to a single certificat using `CERT_PATH` argument.
`make fn-nebula-cert-check-test`: Let the unit tests run!
`fn-nebula-cert-check-test-coverage`: Measure the test coverage... ;-)