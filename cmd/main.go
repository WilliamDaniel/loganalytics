package main

import (
	"fmt"

	"github.com/WilliamDaniel/loganalytics/services/logparser"
	"github.com/WilliamDaniel/loganalytics/services/logreader"
	"github.com/WilliamDaniel/loganalytics/services/logreader/impllogreader"
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
	_, err = parserService.Parse()
	if err != nil {
		panic(err)
	}

	// TODO: obter resultado do parser e enviar para serviço de storage

	// TODO: Chamar serviço process e export em csv.
}
