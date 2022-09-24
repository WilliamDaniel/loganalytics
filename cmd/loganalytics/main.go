package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/WilliamDaniel/loganalytics/services/logstorer"
)

func main() {
	logReaderService := getLogReaderService()
	logFileReader, err := logReaderService.ReadFile()
	if err != nil {
		fmt.Println("error to read log file")
	}

	parserService := getLogParserService(logFileReader)
	parsedLogs, err := parserService.Parse()
	if err != nil {
		log.Fatal(err)
	}

	storageService := getLogStorer()
	logData := logstorer.LogData{
		Logs: *parsedLogs,
	}
	err = storageService.Insert(logData)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("server started")
	log.Fatal(http.ListenAndServe(config.HTTP.Port, nil))
	// TODO: Chamar serviço process e export em csv.
}
