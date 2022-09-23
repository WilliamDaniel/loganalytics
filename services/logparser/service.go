package logparser

import (
	"encoding/json"

	"github.com/WilliamDaniel/loganalytics/shared"
)

type Service interface {
	Parse() (*[]shared.ParsedLog, error)
}
type service struct {
	LogFile *shared.LogFile
}

func NewService(LogFile *shared.LogFile) Service {
	return &service{
		LogFile: LogFile,
	}
}

func (s *service) Parse() (*[]shared.ParsedLog, error) {
	parsedLogLines, err := getParsedLogLines(s.LogFile)
	if err != nil {
		return nil, err
	}
	return &parsedLogLines, nil
}

func getParsedLogLines(LogFile *shared.LogFile) ([]shared.ParsedLog, error) {
	if len(LogFile.Content) == 0 {
		return []shared.ParsedLog{}, nil
	}

	var logs []shared.ParsedLog
	for _, line := range LogFile.Content {
		var parsed shared.ParsedLog

		err := json.Unmarshal([]byte(line), &parsed)
		if err != nil {
			return []shared.ParsedLog{}, err
		}

		logs = append(logs, parsed)
	}
	return logs, nil
}
