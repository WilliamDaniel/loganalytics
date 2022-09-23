package logreader

import (
	"bufio"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCheckLogFile(t *testing.T) {
	const path = "testdata"

	tests := []struct {
		name          string
		logFile       LogFile
		expectedError string
	}{
		{
			name: "Load a log file with 2 lines",
			logFile: LogFile{
				Content: loadFile(filepath.Join(path, "logs_test.txt")),
			},
			expectedError: "",
		},
		{
			name: "Load a empty log file",
			logFile: LogFile{
				Content: loadFile(filepath.Join(path, "empty_log.txt")),
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

func loadFile(filepath string) []string {
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
