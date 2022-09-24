package impllogstorer

import (
	"github.com/WilliamDaniel/loganalytics/services/logstorer"
)

type MySQLDbAdapter struct {
}

func NewMySQLDbAdapter() logstorer.LogRepository {
	return &MySQLDbAdapter{}
}

func (memory *MySQLDbAdapter) Store(l logstorer.LogData) error {

	return nil
}

func (memory *MySQLDbAdapter) Find(ServiceID string) *logstorer.LogData {

	return nil
}
