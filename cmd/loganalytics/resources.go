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

func getLogReaderService(sourcepath string) logreader.Service {
	if logReaderService == nil {
		logReaderService = logreader.NewService(impllogreader.NewLogReaderAdapter(sourcepath))
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
	if logStorerService == nil {
		logStorerService = logstorer.NewService(impllogstorer.NewMemoryDbAdapter(shared.NewMemoryDb()))
	}
	return logStorerService
}
