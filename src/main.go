package main

import (
	"flag"
	"tictactoe-service/database"
	server "tictactoe-service/server"
	"tictactoe-service/util"
)

func main() {
	initFlags()

	server.InitServer()
	database.InitDB()
	server.Serve()
}

func initFlags() {
	emptyFlag := flag.Bool("empty", false, "Enable empty mode")

	flag.Parse()

	if *emptyFlag {
		util.SetupEnvironmentVariables("properties/empty.properties")
	} else {
		util.SetupEnvironmentVariables("properties/default.properties")
	}
}
