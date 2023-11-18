package main

import (
	"tictactoe-service/database"
	server "tictactoe-service/server"
)

func main() {

	server.InitServer()
	database.InitDB()
	server.Serve()
	database.CloseDB()

}
