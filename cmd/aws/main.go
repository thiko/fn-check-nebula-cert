package main

import (
	"context"
	"encoding/json"
	"log"
	"os"
	"thiko/function/check-nebula-cert/pkg/common"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

//Record each data record
type Record struct {
	EventSource    string
	EventSourceArn string
	AWSRegion      string
	S3             events.S3Entity
	SQS            events.SQSMessage // TODO: probably not needed
	SNS            events.SNSEntity  // TODO: probably not needed
}

//Event incoming event
type Event struct {
	Records []Record
}

func HandleLambdaRequest(ctx context.Context, event Event) {

	eventJson, _ := json.MarshalIndent(event, "", "  ")
	log.Printf("EVENT: %s", eventJson)

	runConfig := &common.RunConfig{
		Repository:            &AwsRepository{},
		Mode:                  "aws",
		CertificateBucketName: os.Getenv("cert-bucket"),
		ResultBucketName:      os.Getenv("result-bucket"),
	}

	worker := common.DefaultWorker{}
	worker.Work(runConfig)
}

func main() {
	lambda.Start(HandleLambdaRequest)
}
