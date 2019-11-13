package main

import (
	"context"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/kawakattsun/face-auth-api/internal/collectface"
)

func handler(ctx context.Context, event events.S3Event) (interface{}, error) {
	var bucket string
	var key string
	for _, record := range event.Records {
		bucket = record.S3.Bucket.Name
		key = record.S3.Object.Key
	}
	return collectface.Handler(bucket, key), nil
}

func main() {
	log.Print("Start collectface lambda function.\n")
	lambda.Start(handler)
}
