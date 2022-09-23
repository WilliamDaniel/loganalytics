package logparser

import (
	"encoding/json"

	"github.com/WilliamDaniel/loganalytics/shared"
)

type Service interface {
	Parse() (*[]ParsedLog, error)
}
type service struct {
	LogFile *shared.LogFile
}

func NewService(LogFile *shared.LogFile) Service {
	return &service{
		LogFile: LogFile,
	}
}

func (s *service) Parse() (*[]ParsedLog, error) {
	parsedLogLines, err := getParsedLogLines(s.LogFile)
	if err != nil {
		return nil, err
	}
	return &parsedLogLines, nil
}

func getParsedLogLines(LogFile *shared.LogFile) ([]ParsedLog, error) {
	if len(LogFile.Content) == 0 {
		return []ParsedLog{}, nil
	}

	var logs []ParsedLog
	for _, line := range LogFile.Content {
		var parsed ParsedLog

		err := json.Unmarshal([]byte(line), &parsed)
		if err != nil {
			return []ParsedLog{}, err
		}

		logs = append(logs, parsed)
	}
	return logs, nil
}
