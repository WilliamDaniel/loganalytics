package main

import (
	"log"

	"github.com/crgimenes/goconfig"
)

func init() {
	err := goconfig.Parse(&config)
	if err != nil {
		log.Fatal(err)
	}
}
