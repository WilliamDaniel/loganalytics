package logreader

type Service interface {
	ReadFile() (*LogFile, error)
}

type service struct {
	logReader LogReaderGateway
}

func NewService(logReaderGateway LogReaderGateway) Service {
	return service{
		logReader: logReaderGateway,
	}
}

func (s service) ReadFile() (*LogFile, error) {
	logFile, err := s.logReader.Read()
	if err != nil {
		return nil, ErrToReadFile
	}
	return logFile, nil
}
