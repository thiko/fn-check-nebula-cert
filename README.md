# About
What is nebula? Overlay network developed by Slack. See: https://github.com/slackhq/nebula

# Check Nebula Cert Function

- Reads all .crt files from a defined Bucket
- Validates each .crt file including expiration check
- Stores the result in a file and stores that file in the same bucket. Its using the same file on each invokation (overwrite it).


## Execution

Given Makefile supports different execution & testing modes.

`make exec-dev`: Runs the application in local mode using a dummy bucket.
`make exec-test`: Let the unit tests run!
`make exec-test-coverage`: Measure the test coverage... ;-)