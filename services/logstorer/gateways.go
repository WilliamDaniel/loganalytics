package logstorer

type LogRepository interface {
	Store(LogData) error
}
