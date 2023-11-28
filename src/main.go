package main

import (
	"flag"
	"tictactoe-service/database"
	server "tictactoe-service/server"
	"tictactoe-service/util"
)

var err error

func main() {
	doneParsingFlags := make(chan bool)
	go initFlags(doneParsingFlags)
	<-doneParsingFlags

	server.InitServer()
	database.InitDB()
	server.Serve()
}

func initFlags(done chan bool) {
	emptyFlag := flag.Bool("empty", false, "Enable empty mode")

	flag.Parse()

	if *emptyFlag {
		util.SetupEnvironmentVariablesChan("properties/empty.properties", done)
	} else {
		util.SetupEnvironmentVariablesChan("properties/default.properties", done)
	}
}
