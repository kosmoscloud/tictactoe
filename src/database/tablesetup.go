package database

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func setupUserTable(DB *sql.DB) {
	// there must be some mapper to map the user struct to the query string
	_, err := DB.Exec("CREATE TABLE IF NOT EXISTS users (id INT AUTO_INCREMENT PRIMARY KEY, username VARCHAR(255), created TIMESTAMP)")
	if err != nil {
		log.Fatal(err)
	} else {
		log.Default().Print("Users table setup successfully!")
	}
}

func setupRoomTable(DB *sql.DB) {
	_, err := DB.Exec("CREATE TABLE IF NOT EXISTS rooms (id INT AUTO_INCREMENT PRIMARY KEY, created TIMESTAMP, user1 INT, user2 INT, winner INT DEFAULT 0, moves VARCHAR(255) DEFAULT '0')")
	if err != nil {
		log.Fatal(err)
	} else {
		log.Default().Print("Rooms table setup successfully!")
	}
}
