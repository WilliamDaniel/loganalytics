package impllogstorer

import (
	"errors"
	"fmt"

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
	if len(l.Logs) == 0 {
		return errors.New("nothing to store")
	}

	for _, log := range l.Logs {
		memory.db[log.Service.ID] = log
		fmt.Println("log inserted to service ", log.Service.ID)
	}

	return nil
}

func (memory *MemoryDbAdapter) Find(ServiceID string) *logstorer.LogData {
	if log, ok := memory.db[ServiceID]; ok {
		return &logstorer.LogData{
			Logs: []shared.ParsedLog{
				log,
			},
		}
	}
	return nil
}
