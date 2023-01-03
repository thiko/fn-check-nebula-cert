# About
What is nebula? Overlay network developed by Slack. See: https://github.com/slackhq/nebula


### Nice to know...

- Interesting article: https://www.crimsonmacaw.com/blog/handling-multiple-aws-lambda-event-types-with-go/
- Blank example by AWS: https://github.com/awsdocs/aws-lambda-developer-guide/tree/main/sample-apps/blank-go
- Cheat sheet: https://theburningmonk.com/cloudformation-ref-and-getatt-cheatsheet/
- Good CF recipes with lambda: https://octopus.com/blog/deploying-lambda-cloudformation
- Again reciped: https://octopus.com/blog/tag/CloudFormation

- Structure: https://stackoverflow.com/questions/50904560/how-to-structure-go-application-to-produce-multiple-binaries + https://ieftimov.com/posts/golang-package-multiple-binaries/

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

funknest-certificate-bucket-000001
funknest-lambda-bucket-000001
funknest-s3-object-bucket-000001
funknest-certificate-result-bucket-000001
https://us-west-1.console.aws.amazon.com/cloudformation/home?region=us-west-1#/stacks/quickcreate?templateURL=https%3A%2F%2Fs3.us-west-1.amazonaws.com%2Fcf-templates-es49kqfxr78e-us-west-1%2F2022-12-11T103614.798Zobz-0_aws_cf_setup_s3.yml&stackName=CertificateValidationS3Buckets&param_CertificateBucketName=funknest-certificate-bucket-000001&param_ResultBucketName=funknest-certificate-result-bucket-000001&param_LambdaS3Object=funknest-s3-object-bucket-000001&param_LambdaS3Bucket=funknest-lambda-bucket-000001

FIXME: https://stackoverflow.com/questions/61845013/package-xxx-is-not-in-goroot-when-building-a-go-project