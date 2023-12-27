package database

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func InitDB() {
	var err error
	DB, err = sql.Open("mysql", "root:password@tcp(127.0.0.1:3306)/tictactoe_database?parseTime=true")
	if err != nil {
		log.Fatal(err)
	} else {
		log.Default().Print("Successfully connected to tictactoe-database!")
	}

	SetupUserTable()
	SetupRoomTable()
	SetupMoveTable()
}

func CloseDB() {
	DB.Close()
}
