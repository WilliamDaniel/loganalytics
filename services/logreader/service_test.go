package logreader

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type logReaderAdapterMock struct {
	filepath string
}

func newLogReaderAdapterMock(filepath string) logReaderAdapterMock {
	return logReaderAdapterMock{
		filepath: filepath,
	}
}

func Test_service_ReadFile(t *testing.T) {
	tests := []struct {
		name            string
		logReader       LogReaderGateway
		expectedLogFile *LogFile
		expectedError   string
	}{
		{
			name:      "empty file",
			logReader: newLogReaderAdapterMock("empty_file.txt"),
			expectedLogFile: &LogFile{
				Content: nil,
			},
			expectedError: "",
		},
		{
			name:      "with logs",
			logReader: newLogReaderAdapterMock("logs.txt"),
			expectedLogFile: &LogFile{
				Content: []string{"{'request:test'}"},
			},
			expectedError: "",
		},
		{
			name:      "with bad filepath",
			logReader: newLogReaderAdapterMock("logss.txt"),
			expectedLogFile: &LogFile{
				Content: []string{"{'request:test'}"},
			},
			expectedError: ErrToReadFile.Error(),
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			s := service{
				logReader: test.logReader,
			}
			logFile, err := s.ReadFile()
			if test.expectedError != "" {
				assert.EqualError(t, err, test.expectedError)
			} else {
				assert.NotNil(t, logFile)
				assert.NoError(t, err)
			}
		})
	}
}

func (r logReaderAdapterMock) Read() (*LogFile, error) {
	var logFile LogFile
	switch r.filepath {
	case "empty_file.txt":
		logFile = LogFile{
			Content: nil,
		}
	case "logs.txt":
		logFile = LogFile{
			Content: []string{"{'request:test'}"},
		}
	default:
		return nil, ErrToReadFile
	}

	return &logFile, nil
}
