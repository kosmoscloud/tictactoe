package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

var (
	DB  *sql.DB
	err error
)

var fatalExit = log.Fatal
var sqlOpenString = generateSqlOpenString

// InitDB initializes the database connection and uses the fatalExit function if an error occurs
func InitDB() {
	openSqlConnection()
	pingDatabase()
	setupUserTable()
}

func openSqlConnection() {
	DB, err = sql.Open("mysql", sqlOpenString())
	if err != nil {
		fatalExit(err)
	}
}

func pingDatabase() {
	err = DB.Ping()
	if err != nil {
		fatalExit(err)
	} else {
		log.Println("Successfully pinged tictactoe-database!")
	}
}

func generateSqlOpenString() string {
	username := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	database := os.Getenv("DB_NAME")

	sqlOpenString :=
		fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", username, password, host, port, database)

	return sqlOpenString
}
