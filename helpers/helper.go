package helpers

import (
	"log"

	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/s3music-to-dynamo-record/helpers/s3helpers"
)

func GetSongFromS3(key string) error{
	s3Session := s3helpers.GetSession()
	getObjectInput := s3.GetObjectInput{
		Bucket: &s3helpers.S3MusicBucket,
		Key: &key,
	}

	output, err := s3Session.GetObject(&getObjectInput)
	if err != nil {
		return err
	}
	log.Println(output.Body)
	return nil
}