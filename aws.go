package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	"github.com/aws/aws-lambda-go/events"
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

	work(runConfig{
		repository:            &AwsRepository{},
		mode:                  "aws",
		certificateBucketName: os.Getenv("cert-bucket"),
		resultBucketName:      os.Getenv("result-bucket"),
	})
}
