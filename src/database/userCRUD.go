package database

import (
	"tictactoe-service/server/dto"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func GetUser(id int64) (*dto.User, error) {
	user := &dto.User{}
	row := DB.QueryRow("SELECT * FROM users WHERE id=?", id)
	err := row.Scan(&user.UserId, &user.Username, &user.CreatedDate)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func CreateUser(username string) (*dto.User, error) {
	createdDate := time.Now()
	rows, err := DB.Exec("INSERT INTO users (username, created) VALUES (?, ?)", username, createdDate)
	if err != nil {
		return nil, err
	}
	id, err := rows.LastInsertId()
	if err != nil {
		return nil, err
	}

	user, err := GetUser(id)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func UpdateUser(id int64, username string) (*dto.User, error) {
	_, err := DB.Exec("UPDATE users SET username=? WHERE id=?", username, id)
	if err != nil {
		return nil, err
	}

	user, err := GetUser(id)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func DeleteUser(id int64) (*dto.User, error) {
	_, err := DB.Exec("DELETE FROM users WHERE id=?", id)
	if err != nil {
		return nil, err
	}
	return nil, nil
}
