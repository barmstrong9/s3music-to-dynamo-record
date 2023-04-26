package s3helpers

import (
	"bytes"
	"io/ioutil"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3iface"
	"github.com/s3music-to-dynamo-record/util"
)
var (
	s3Session	s3iface.S3API
	awsRegion	= util.GetEnv("AWS_DEFAULT_REGION", "eu-west-1")
	S3MusicBucket	= util.GetEnv("redheadrhythms-music", "redheadrhythms-music")
)

type mockBucket map[string][]byte

type mockS3API struct {
	s3iface.S3API
	mockBucket map[string]mockBucket
}

//CreateSession creates a new S3 session 
func CreateSession() {
	awsConfig := aws.Config{Region: aws.String(awsRegion)}
	session := session.Must(session.NewSession(&awsConfig))
	s3Session = s3.New(session)
}

func GetSession() s3iface.S3API {
	return s3Session
}

func CreateMockSession() {
	s3Session = &mockS3API{
		S3API: nil,
		mockBucket:  map[string]mockBucket{S3MusicBucket: {"01_Robot_Rock_Oh_Yeah.mp3": []byte("some mp3 data")}},
	}
}


func (m *mockS3API) GetObject(input *s3.GetObjectInput) (*s3.GetObjectOutput, error) {
	if bucket, ok := m.mockBucket[*input.Bucket]; ok {
		if obj, ok := bucket[*input.Key]; ok {
			return &s3.GetObjectOutput{
				Body: ioutil.NopCloser(bytes.NewReader(obj)),
			}, nil
		}
	}
	return nil, awserr.New(s3.ErrCodeNoSuchKey, "The specified key does not exist.", nil)
}
