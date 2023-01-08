#!/bin/bash
set -e
set -o pipefail

STACK_NAME=nebula-cert-check-buckets

RANDOM_ID=$(dd if=/dev/random bs=8 count=1 2>/dev/null | od -An -tx1 | tr -d ' \t\n')
CF_TEMPLATE_BUCKET=nebula-cert-cf-bucket-$RANDOM_ID

LAMBDA_BUCKET=test-lambda-bucket-000991

# build the lambda function
make build-lambda
echo "lambda function build completed"

# create S3 buckets
aws cloudformation delete-stack --stack-name $STACK_NAME

echo "Create bucket for template uploads: $CF_TEMPLATE_BUCKET"
aws s3 mb s3://$CF_TEMPLATE_BUCKET


aws cloudformation package \
    --template-file ./cmd/aws/infrastructure/0_aws_cf_setup_s3.yml \
    --s3-bucket $CF_TEMPLATE_BUCKET \
    --output-template-file nebula-cert-check-buckets-out.yml

aws cloudformation deploy \
 --template-file nebula-cert-check-buckets-out.yml \
 --stack-name $STACK_NAME \
 --parameter-overrides \
    CertificateBucketName="test-cert-bucket-000991" \
    ResultBucketName="test-result-bucket-000991" \
    LambdaS3Bucket=$LAMBDA_BUCKET \
 --capabilities CAPABILITY_NAMED_IAM
echo "buckets created"

# upload lambda zip file
aws s3 cp ./bin/aws_lambda.zip s3://$LAMBDA_BUCKET/
echo "lambda function uploaded to s3 bucket"

# create lambda function
aws cloudformation package \
    --template-file ./cmd/aws/infrastructure/1_aws_cf_setup_lambda.yml \
    --output-template-file nebula-cert-check-lambda-out.yml

aws cloudformation deploy \
    --template-file nebula-cert-check-lambda-out.yml \
    --stack-name nebula-cert-check-lambda \
    --capabilities CAPABILITY_NAMED_IAM 
    
echo "lambda function created on aws"
