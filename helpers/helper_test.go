package helpers

import (
	"flag"
	"os"
	"testing"

	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/s3music-to-dynamo-record/helpers/s3helpers"
	"github.com/stretchr/testify/assert"
)


func TestMain(m *testing.M) {
	flag.Parse()

	s3helpers.CreateMockSession()

	exitCode := m.Run()

	os.Exit(exitCode)
}

func TestGetSongFromS3(t *testing.T) {
	tests := []struct{
		input string
		expected error
	}{
		{
			"01_Robot_Rock_Oh_Yeah.mp3",
			nil,
		},
		{
			"",
			awserr.New(s3.ErrCodeNoSuchKey, "The specified key does not exist.", nil),
		},
	}

	for _, test := range tests {
		_, err := GetSongFromS3(test.input)
		assert.Equal(t, test.expected, err)
	}
}
