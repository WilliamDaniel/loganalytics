package logreader

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCheckLogFile(t *testing.T) {

	tests := []struct {
		name          string
		logFile       LogFile
		expectedError string
	}{
		{
			name: "Load a log file with 2 lines",
			logFile: LogFile{
				Content: []string{"{'request':'test'}"},
			},
			expectedError: "",
		},
		{
			name: "Load a empty log file",
			logFile: LogFile{
				Content: nil,
			},
			expectedError: ErrEmptyLogFile.Error(),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			err := test.logFile.Check()

			if test.expectedError != "" {
				assert.EqualError(t, err, test.expectedError)
			} else {
				assert.NoError(t, err, "unexpected error")
			}
		})
	}
}
