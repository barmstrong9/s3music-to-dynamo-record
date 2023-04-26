package helpers

import (
	"log"

	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/s3music-to-dynamo-record/helpers/s3helpers"
)

type SongMetadata struct {
	Author   string
	Duration int
	Genre    string
}

func GetSongFromS3(key string) error {
	s3Session := s3helpers.GetSession()
	headObjectInput := s3.HeadObjectInput{
		Bucket: &s3helpers.S3MusicBucket,
		Key:    &key,
	}
	output, err := s3Session.HeadObject(&headObjectInput)
	if err != nil {
		return nil
	}
	log.Println(output.Metadata, "RIGHT HERE")

	return nil
}
