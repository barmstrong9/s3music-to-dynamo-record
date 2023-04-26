package helpers

import (
	"errors"
	"log"
	"strconv"

	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/s3music-to-dynamo-record/helpers/s3helpers"
)

type SongMetadata struct {
	Author string
	Duration int
	Genre string
}

func GetSongFromS3(key string) (s3.GetObjectOutput, error){
	s3Session := s3helpers.GetSession()
	getObjectInput := s3.GetObjectInput{
		Bucket: &s3helpers.S3MusicBucket,
		Key: &key,
	}

	output, err := s3Session.GetObject(&getObjectInput)
	if err != nil {
		return s3.GetObjectOutput{}, err
	}
	return *output, nil
}

func GetMetadata(s3Output *s3.GetObjectOutput) (string, int, string, error) {
	log.Println(s3Output.Metadata)
	author, ok1 := s3Output.Metadata["Author"]
	durationStr, ok2 := s3Output.Metadata["Duration"]
	genre, ok3 := s3Output.Metadata["Genre"]
	if !ok1 || !ok2 || !ok3 {
		return "", 0, "", errors.New("metadata not found in s3 object")
	}

	duration, err := strconv.Atoi(*durationStr)
	if err != nil {
		return "", 0, "", errors.New("failed to parse duration string")
	}

	return *author, duration, *genre, nil
}