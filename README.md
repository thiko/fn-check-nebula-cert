# About
What is nebula? Overlay network developed by Slack. See: https://github.com/slackhq/nebula


### Nice to know...

- Interesting article: https://www.crimsonmacaw.com/blog/handling-multiple-aws-lambda-event-types-with-go/
- Blank example by AWS: https://github.com/awsdocs/aws-lambda-developer-guide/tree/main/sample-apps/blank-go

# Check Nebula Cert Function

- Reads all .crt files from a defined Bucket
- Validates each .crt file including expiration check
- Stores the result in a file and stores that file in the same bucket. Its using the same file on each invokation (overwrite it).


## Execution

Given Makefile supports different execution & testing modes.

- `make exec-dev`: Runs the application in local mode using a dummy bucket.
- `make exec-test`: Let the unit tests run!
- `make exec-test-coverage`: Measure the test coverage... ;-)


## Deployment

### AWS
**Validation**: `aws cloudformation validate-template --region us-east-1 --template-body file://infrastructure/aws_cf_setup.yml`


