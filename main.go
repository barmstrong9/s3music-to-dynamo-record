package main

import (
	"log"
	"context"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-lambda-go/events"
)

func main() {
	lambda.Start(Handler)
}

func Handler(ctx context.Context, s3Event events.S3Event) {
	for _, record := range s3Event.Records {
		s3 := record.S3
		log.Printf("[%s - %s] Bucket = %s, Key = %s \n", record.EventSource, record.EventTime, s3.Bucket.Name, s3.Object.Key)
}
}

//Remember to do
// GOOS=linux go build -o build/main cmd/main.go
// zip -jrm build/main.zip build/main
// Then upload to s3 if CI/CD has not been set up.