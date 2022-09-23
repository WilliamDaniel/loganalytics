package logstorer

import "errors"

type Service interface {
	Insert(LogData) error
}

type service struct {
	repo LogRepository
}

func NewService(repo LogRepository) Service {
	return &service{
		repo: repo,
	}
}

func (s *service) Insert(l LogData) error {
	return errors.New("not implemented yet")
}
