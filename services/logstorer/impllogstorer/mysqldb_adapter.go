package impllogstorer

import (
	"github.com/WilliamDaniel/loganalytics/services/logstorer"
	"github.com/jinzhu/gorm"
)

type MySQLDbAdapter struct {
	db *gorm.DB
}

func NewMySQLDbAdapter(dataBase *gorm.DB) logstorer.LogRepository {
	return &MySQLDbAdapter{
		db: dataBase,
	}
}

func (mysql *MySQLDbAdapter) Store(logs logstorer.LogData) error {
	mysql.db.Create(&logs)
	return nil
}

func (mysql *MySQLDbAdapter) Find(ServiceID string) *logstorer.LogData {

	return nil
}
