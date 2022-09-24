package main

var config struct {
	HTTP struct {
		Port string `cfgDefault:":8900"`
	}
	Log struct {
		Filepath string `cfgDefault:"blackhole/logs.txt"`
	}
	DB struct {
		Type string `cfgDefault:"MEMORY"` // MEMORY, MYSQL
	}
}
