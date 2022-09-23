package logreader

import "github.com/WilliamDaniel/loganalytics/shared"

type Service interface {
	ReadFile() (*shared.LogFile, error)
}

type service struct {
	logReader LogReaderGateway
}

func NewService(logReaderGateway LogReaderGateway) Service {
	return service{
		logReader: logReaderGateway,
	}
}

func (s service) ReadFile() (*shared.LogFile, error) {
	logFile, err := s.logReader.Read()
	if err != nil {
		return nil, ErrToReadFile
	}
	return logFile, nil
}
