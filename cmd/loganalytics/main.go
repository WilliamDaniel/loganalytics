package main

import (
	"fmt"

	"github.com/WilliamDaniel/loganalytics/services/logstorer"
)

func main() {
	logReaderService := getLogReaderService("services/logreader/testdata/logs.txt")
	logFileReader, err := logReaderService.ReadFile()
	if err != nil {
		fmt.Println("error to read log file")
	}

	parserService := getLogParserService(logFileReader)
	parsedLogs, err := parserService.Parse()
	if err != nil {
		panic(err)
	}

	storageService := getLogStorer()
	logData := logstorer.LogData{
		Logs: *parsedLogs,
	}
	err = storageService.Insert(logData)
	if err != nil {
		fmt.Println(err)
	}

	// TODO: Chamar serviço process e export em csv.
}
