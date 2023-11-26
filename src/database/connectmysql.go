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

func InitDB() error {
	DB, err = sql.Open("mysql", generateSqlOpenString())
	if err != nil {
		log.Println(err)
		return err
	}

	err = DB.Ping()
	if err != nil {
		log.Println(err)
		return err
	} else {
		log.Println("Successfully connected to tictactoe-database!")
	}

	setupUserTable(DB)
	return nil
}

func CloseDB() error {
	err = DB.Close()
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
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
