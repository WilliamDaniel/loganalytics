package main

import (
	"fmt"

	"github.com/WilliamDaniel/loganalytics/services/logparser"
	"github.com/WilliamDaniel/loganalytics/services/logreader"
	"github.com/WilliamDaniel/loganalytics/services/logreader/impllogreader"
	"github.com/WilliamDaniel/loganalytics/services/logstorer"
	"github.com/WilliamDaniel/loganalytics/services/logstorer/impllogstorer"
	"github.com/WilliamDaniel/loganalytics/shared"
)

func main() {
	logReaderAdapter := impllogreader.NewLogReaderAdapter("services/logreader/testdata/logs_test.txt")
	logReaderService := logreader.NewService(logReaderAdapter)
	logFileReader, err := logReaderService.ReadFile()
	if err != nil {
		fmt.Println("error to read log file")
	}

	fmt.Print(logFileReader)

	parserService := logparser.NewService(logFileReader)
	parsedLogs, err := parserService.Parse()
	if err != nil {
		panic(err)
	}

	storageAdapter := impllogstorer.NewMemoryDbAdapter(shared.NewMemoryDb())
	storageService := logstorer.NewService(storageAdapter)
	logData := logstorer.LogData{
		Log: *parsedLogs,
	}
	err = storageService.Insert(logData)
	if err != nil {
		panic(err)
	}

	// TODO: Chamar serviço process e export em csv.
}
