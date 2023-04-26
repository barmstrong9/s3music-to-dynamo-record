package main

import (
	"errors"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/s3music-to-dynamo-record/helpers"
	"github.com/s3music-to-dynamo-record/helpers/s3helpers"
)

func main() {
	s3helpers.CreateSession()
	lambda.Start(handler)
}

func handler(s3Event events.S3Event) error{
	if len(s3Event.Records) == 0 {
		return errors.New("invalid response, S3 Event is empty")
	}
	for _, record := range s3Event.Records {
		s3Record := record.S3
		log.Printf("[%s - %s] Bucket = %s, Key = %s \n", record.EventSource, record.EventTime, s3Record.Bucket.Name, s3Record.Object.Key)
		objectOutput, err := helpers.GetSongFromS3(s3Record.Object.Key)
		if err != nil {
			return err
		}
		artist, duration, genre, err := helpers.GetMetadata(&objectOutput)
		if err != nil {
			return err
		}

		log.Println(artist, duration, genre, "HERE MAN")
	}
	return nil
}
