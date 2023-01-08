# About
What is nebula? Overlay network developed by Slack. See: https://github.com/slackhq/nebula

## Current state
- CF Files works when doing it manually. Bash scripts has still some issues. 
- Havent tried the lambda in action so far ;)

## Process

- Reads all .crt files from a defined Bucket
- Validates each .crt file including expiration check
- Stores the result in a file and stores that file in the same bucket. Its using the same file on each invokation (overwrite it).


# Project structure

```
fn-check-nebula-cert
├── bin
├── cmd
│   ├── aws
│   │   ├── aws_repository.go
│   │   ├── infrastructure
│   │   │   ├── 0_aws_cf_setup_s3.yml
│   │   │   └── 1_aws_cf_setup_lambda.yml
│   │   └── main.go
│   ├── gcp
│   │   └── gcp_repository.go
│   └── local
│       ├── local.go
│       └── local_repository.go
├── go.mod
├── go.sum
├── LICENSE
├── Makefile
├── pkg
│   └── common
│       ├── common.go
│       ├── worker.go
│       └── worker_test.go
└── README.md
```

`1_aws_cf_setup_lambda.yml`: Reads most parameter out of Systems manager parameter store. Does always use the most recent version of the parameter.


## Nice to know...

- Interesting article: https://www.crimsonmacaw.com/blog/handling-multiple-aws-lambda-event-types-with-go/
- Blank example by AWS: https://github.com/awsdocs/aws-lambda-developer-guide/tree/main/sample-apps/blank-go
- Cheat sheet: https://theburningmonk.com/cloudformation-ref-and-getatt-cheatsheet/
- Good CF recipes with lambda: https://octopus.com/blog/deploying-lambda-cloudformation
- Again reciped: https://octopus.com/blog/tag/CloudFormation

- Structure: https://stackoverflow.com/questions/50904560/how-to-structure-go-application-to-produce-multiple-binaries + https://ieftimov.com/posts/golang-package-multiple-binaries/



## Execution

Given Makefile supports different execution & testing modes.

- `build-lambda`: build AWS Lambda zip file. Obviously this only includes `cmd/aws` parts.
- `make exec-dev`: Runs the application in local mode using a dummy bucket.
- `make exec-test`: Let the unit tests run!
- `make exec-test-coverage`: Measure the test coverage... ;-)


## Deployment

### AWS
**Validation**: `aws cloudformation validate-template --region us-east-1 --template-body file://infrastructure/aws_cf_setup.yml`

**First** execute the `0_aws_cf_setup_s3.yml` CloudFormation template. Either by uploading it manually or by using following link: [AWS CF Creation](https://us-west-1.console.aws.amazon.com/cloudformation/home?region=us-west-1#/stacks/quickcreate?templateURL=https%3A%2F%2Fs3.us-west-1.amazonaws.com%2Fcf-templates-es49kqfxr78e-us-west-1%2F2023-01-07T101731.840Zb9h-0_aws_cf_setup_s3.yml&stackName=NebulaCertCheckS3Buckets&param_CertificateBucketName=test-cert-bucket-000991&param_ResultBucketName=test-result-bucket-000991&param_LambdaS3Bucket=test-lambda-bucket-000991)
**Change the S3 Bucket names**. Buckets names are globally unique.

**Second** build the lambda function: `make build-lambda`. 

**Third** Upload the lambda zip file to the S3 bucket(LambdaS3Bucket).

**Fourth** Execute the second CloudFormation template `1_aws_cf_setup_lambda.yml`

Thats it :-)

#### Open points

- AWS: Direct upload of go binary (https://catalog.workshops.aws/cfn101/en-US/intermediate/templates/package-and-deploy)

- Bucket policies
```
  CertificateBucketPolicy:
    Type: AWS::S3::BucketPolicy
    Metadata:
      comment: 'Bucket policy to allow all users get access but deny put access.'
    Properties:
      Bucket: !Ref CertificationBucket
      PolicyDocument:
        Statement:
        - Action: ['s3:GetObject', 's3:ListBucket']
          Effect: Allow
          Principal: '*'
          Resource:
          - !Sub arn:aws:s3:::${CertificationBucket}
          - !Sub arn:aws:s3:::${CertificationBucket}/*

  ResultBucketPolicy:
    Type: AWS::S3::BucketPolicy
    Metadata:
      comment: ''
    Properties:
      Bucket: !Ref ResultBucket
      PolicyDocument:
        Statement:
        - Action: ['s3:GetObject', 's3:ListBucket', 's3:PutObject']
          Effect: Allow
          Principal: '*'
          Resource:
          - !Sub arn:aws:s3:::${ResultBucket}
          - !Sub arn:aws:s3:::${ResultBucket}
```