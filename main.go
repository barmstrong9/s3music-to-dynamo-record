package main

import (
	"errors"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/s3music-to-dynamo-record/helpers/s3helpers"
)

func main() {
	s3helpers.CreateSession()
	lambda.Start(handler)
}

func handler(s3Event events.S3Event) error{
	s3Session := s3helpers.GetSession()
	if len(s3Event.Records) == 0 {
		return errors.New("invalid response, S3 Event is empty")
	}
	for _, record := range s3Event.Records {
		s3Record := record.S3
		log.Printf("[%s - %s] Bucket = %s, Key = %s \n", record.EventSource, record.EventTime, s3Record.Bucket.Name, s3Record.Object.Key)
		input := s3.GetObjectInput{
			Bucket: &s3Record.Bucket.Name,
			Key: &s3Record.Object.Key,
		}

		output, err := s3Session.GetObject(&input)
		if err != nil {
			log.Println("Error getting object from S3")
			return err
		}

		log.Println(output)

	}
	return nil
}
