package main

import (
	"log"
	"os"
)

var setPath = logPath()

func gologger() *log.Logger {
	file, err := os.Create(setPath)
	if err != nil {
		log.Fatalln("fail to create dsPing.log file!")
	}
	logger := log.New(file, "[dsPing] ", log.Lshortfile|log.LstdFlags)
	return logger
}
