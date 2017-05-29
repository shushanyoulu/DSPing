package main

import (
	"database/sql"
	"flag"
	"sync"
)

// Init config
var filename = flag.String("f", "../conf/config.json", "JSON configuration file")
var httpPort = flag.Int("p", 8899, "HTTP port")
var lock sync.Mutex

const version = "0.2.3"

var logger = gologger()

// Main function
func main() {
	flag.Parse()
	// Config
	var config Config
	logger.Println("Opening config file: ", *filename)
	config = readConfig(*filename)
	config.Version = version
	logger.Printf("Config loaded")
	logger.Println("LogDB : " + config.Db)
	db, err := sql.Open("sqlite3", config.Db)
	if err != nil {
		logger.Println(err)
	}
	// Running
	res := make(chan TargetStatus)
	state := NewState()
	state.Localname = config.Name
	state.Localip = config.Ip
	for _, target := range config.Targets {
		startPing(db, config, target, res)
	}
	// HTTP
	go startHttp(*httpPort, state, db, config)
	for {
		status := <-res
		state.Lock.Lock()
		state.State[status.Target] = status
		state.Lock.Unlock()
	}

}
