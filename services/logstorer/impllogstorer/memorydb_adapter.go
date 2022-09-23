package impllogstorer

import (
	"errors"

	"github.com/WilliamDaniel/loganalytics/services/logstorer"
	"github.com/WilliamDaniel/loganalytics/shared"
)

type MemoryDbAdapter struct {
	db shared.MemoryDb
}

func NewMemoryDbAdapter(db shared.MemoryDb) logstorer.LogRepository {
	return &MemoryDbAdapter{
		db: db,
	}
}

func (memory *MemoryDbAdapter) Store(l logstorer.LogData) error {
	if len(l.Log) == 0 {
		return errors.New("nothing to store")
	}

	for _, log := range l.Log {
		memory.db[log.Service.ID] = log
	}

	return errors.New("not implemented yet")
}

func (memory *MemoryDbAdapter) Find(ServiceID string) *logstorer.LogData {
	if log, ok := memory.db[ServiceID]; ok {
		return &logstorer.LogData{
			Log: []shared.ParsedLog{
				log,
			},
		}
	}
	return nil
}
