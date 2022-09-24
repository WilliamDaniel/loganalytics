package logstorer

type Service interface {
	Insert(Logdata LogData) error
	Get(ServiceID string) *LogData
}

type service struct {
	repo LogRepository
}

func NewService(repo LogRepository) Service {
	return &service{
		repo: repo,
	}
}

func (s *service) Insert(log LogData) error {
	if err := s.repo.Store(log); err != nil {
		return ErrToStoreLogIntoDatabase
	}
	return nil
}

func (s *service) Get(ServiceID string) *LogData {
	return s.repo.Find(ServiceID)
}
