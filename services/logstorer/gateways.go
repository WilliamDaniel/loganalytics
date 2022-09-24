package logstorer

type LogRepository interface {
	Store(log LogData) error
	Find(ServiceID string) *LogData
}
