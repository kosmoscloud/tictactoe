package database

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func setupUserTable(DB *sql.DB) error {
	_, err := DB.Exec("CREATE TABLE IF NOT EXISTS users (id INT AUTO_INCREMENT PRIMARY KEY, username VARCHAR(255), created TIMESTAMP)")
	if err != nil {
		log.Println(err)
		return err
	} else {
		log.Default().Print("Users table setup successfully!")
		return nil
	}
}
