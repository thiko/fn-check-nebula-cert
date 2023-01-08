#!/bin/bash
set -e
set -o pipefail

STACKS=("nebula-cert-check-lambda" "nebula-cert-check-buckets")


aws s3 rm s3://test-lambda-bucket-000991/aws_lambda.zip
echo "Removed Lambda function zip from bucket"

# if stack names are provided by invoker
if [[ $# -ge 1 ]] ; then
    STACKS=$1
fi
echo "Deleting stacks: ${STACKS[*]}"

for stack in ${STACKS[@]}; do
    FUNCTION=$(aws cloudformation describe-stack-resource --stack-name $stack --logical-resource-id function --query 'StackResourceDetail.PhysicalResourceId' --output text) 
    aws cloudformation delete-stack --stack-name $stack
    echo "Deleted $stack stack."
done


while true; do
    read -p "Delete function log group (/aws/lambda/$FUNCTION)? (y/n)" response
    case $response in
        [Yy]* ) aws logs delete-log-group --log-group-name /aws/lambda/$FUNCTION; break;;
        [Nn]* ) break;;
        * ) echo "Response must start with y or n.";;
    esac
done
