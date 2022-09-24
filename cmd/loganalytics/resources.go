package main

import (
	"github.com/WilliamDaniel/loganalytics/services/logparser"
	"github.com/WilliamDaniel/loganalytics/services/logreader"
	"github.com/WilliamDaniel/loganalytics/services/logreader/impllogreader"
	"github.com/WilliamDaniel/loganalytics/services/logstorer"
	"github.com/WilliamDaniel/loganalytics/services/logstorer/impllogstorer"
	"github.com/WilliamDaniel/loganalytics/shared"
)

var (
	logReaderService logreader.Service
	logParserService logparser.Service
	logStorerService logstorer.Service
)

func getLogReaderService() logreader.Service {
	if logReaderService == nil {
		logReaderService = logreader.NewService(impllogreader.NewLogReaderAdapter(config.Log.Filepath))
	}
	return logReaderService
}

func getLogParserService(log *shared.LogFile) logparser.Service {
	if logParserService == nil {
		logParserService = logparser.NewService(log)
	}
	return logParserService
}

func getLogStorer() logstorer.Service {

	switch config.DB.Type {
	case "MEMORY":
		if logStorerService == nil {
			logStorerService = logstorer.NewService(impllogstorer.NewMemoryDbAdapter(shared.NewMemoryDb()))
		}
	case "MYSQL":
		if logStorerService == nil {
			logStorerService = logstorer.NewService(impllogstorer.NewMySQLDbAdapter(shared.NewMySQLDb().Db))
		}
	default:
		if logStorerService == nil {
			logStorerService = logstorer.NewService(impllogstorer.NewMemoryDbAdapter(shared.NewMemoryDb()))
		}
	}
	return logStorerService
}
