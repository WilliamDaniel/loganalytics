package logparser

import (
	"bufio"
	"os"
	"testing"

	"github.com/WilliamDaniel/loganalytics/shared"
	"github.com/stretchr/testify/assert"
)

func Test_service_Parse(t *testing.T) {

	tests := []struct {
		name               string
		logFile            shared.LogFile
		expectedParsedLogs *[]shared.ParsedLog
		expectedError      string
	}{
		{
			name: "with authenticated entity",
			logFile: shared.LogFile{
				Content: loadFileLines("testdata/logs.txt"),
			},
			expectedParsedLogs: &[]shared.ParsedLog{
				{
					AuthenticatedEntity: shared.RequestAuthenticatedEntity{
						ConsumerID: shared.AuthenticatedEntityConsumerID{
							UUID: "72b34d31-4c14-3bae-9cc6-516a0939c9d6",
						},
					},
				},
			},
		},
		{
			name: "with authenticated entity",
			logFile: shared.LogFile{
				Content: []string{`{"authenticated_entity":{"consumer_id":{"uuid":"123"}}, "service":{"id":"456"}}`},
			},
			expectedParsedLogs: &[]shared.ParsedLog{
				{
					AuthenticatedEntity: shared.RequestAuthenticatedEntity{
						ConsumerID: shared.AuthenticatedEntityConsumerID{
							UUID: "123",
						},
					},
					Service: shared.RequestService{
						ID: "456",
					},
				},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			s := &service{
				LogFile: &test.logFile,
			}
			parsed, err := s.Parse()
			if test.expectedError != "" {
				assert.EqualError(t, err, test.expectedError)
			} else {
				assert.NoError(t, err)
				assert.NotEmpty(t, parsed)
				log := *test.expectedParsedLogs
				logParsed := *parsed

				assert.Equal(t, log[0].AuthenticatedEntity.ConsumerID, logParsed[0].AuthenticatedEntity.ConsumerID)
			}

		})
	}
}

func loadFileLines(filepath string) []string {
	readFile, err := os.Open(filepath)
	if err != nil {
		panic(err)
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	var fileLines []string

	for fileScanner.Scan() {
		fileLines = append(fileLines, fileScanner.Text())
	}
	readFile.Close()

	return fileLines
}
