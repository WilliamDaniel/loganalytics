package impllogreader

import (
	"bufio"
	"os"

	"github.com/WilliamDaniel/loganalytics/shared"
)

type LogReaderAdapter struct {
	Filepath string
}

func NewLogReaderAdapter(Filepath string) LogReaderAdapter {
	return LogReaderAdapter{
		Filepath: Filepath,
	}
}

func (r LogReaderAdapter) Read() (*shared.LogFile, error) {
	fileLines := loadFileLines(r.Filepath)
	logFile := shared.LogFile{
		Content: fileLines,
	}
	return &logFile, nil
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
