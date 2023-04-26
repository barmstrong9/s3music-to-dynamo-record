package main

import (
	"errors"
	"flag"
	"os"
	"testing"

	"github.com/aws/aws-lambda-go/events"
	"github.com/s3music-to-dynamo-record/helpers/s3helpers"
	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	flag.Parse()

	s3helpers.CreateMockSession()

	exitCode := m.Run()

	os.Exit(exitCode)
}

func TestHandler(t *testing.T) {

	tests := []struct {
		input    events.S3Event
		expected error
	}{
		{
			events.S3Event{
				Records: []events.S3EventRecord{
					{
						EventSource: "aws:s3",
						S3: events.S3Entity{
							Bucket: events.S3Bucket{
								Name: "redheadrhythms-music",
							},
							Object: events.S3Object{
								Key: "01_Robot_Rock_Oh_Yeah.mp3",
							},
						},
					},
				},
			},
			nil,
		},
		{
			events.S3Event{
				Records: []events.S3EventRecord{},
			},
			errors.New("invalid response, S3 Event is empty"),
		},
	}

	for _, test := range tests {
		output := handler(test.input)

		assert.Equal(t, test.expected, output)
	}
}
