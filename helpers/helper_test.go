package helpers

import (
	"flag"
	"testing"
	"os"

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
	}

	for _, test := range tests {
		output := GetSongFromS3(test.input)
		assert.Equal(t, test.expected, output)
	}
}
