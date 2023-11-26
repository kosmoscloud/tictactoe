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
	err = database.InitDB()
	if err != nil {
		panic(err)
	}
	server.Serve()
	err = database.CloseDB()
	if err != nil {
		panic(err)
	}
}

func initFlags(done chan bool) {
	testFlag := flag.Bool("t", false, "Enable test mode")

	flag.Parse()

	if *testFlag {
		util.SetupEnvironmentVariables("properties/test.properties", done)
	} else {
		util.SetupEnvironmentVariables("properties/default.properties", done)
	}
}
