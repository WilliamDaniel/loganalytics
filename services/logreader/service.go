package logreader

type Service interface {
	ReadFile(Filepath string) (*LogFile, error)
}

type service struct {
	logReader LogReaderGateway
}

func NewService(logReaderGateway LogReaderGateway) Service {
	return service{
		logReader: logReaderGateway,
	}
}

func (s service) ReadFile(Filepath string) (*LogFile, error) {
	logFile, err := s.logReader.Read()
	if err != nil {
		return nil, ErrToReadFile
	}
	return logFile, nil
}
