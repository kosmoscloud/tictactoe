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
	user, err := GetUser(id)
	if err != nil {
		return nil, err
	}

	_, err = DB.Exec("DELETE FROM users WHERE id=?", id)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func GetRoom(id int64) (*dto.Room, error) {
	room := &dto.Room{}
	row := DB.QueryRow("SELECT * FROM rooms WHERE id=?", id)
	err := row.Scan(&room.RoomId, &room.CreatedDate, &room.User1, &room.User2, &room.Winner, &room.Moves)
	if err != nil {
		return nil, err
	}
	return room, nil
}

func CreateRoom(user1 int64, user2 int64) (*dto.Room, error) {
	createdDate := time.Now()
	rows, err := DB.Exec("INSERT INTO rooms (created, user1, user2) VALUES (?, ?, ?)", createdDate, user1, user2)
	if err != nil {
		return nil, err
	}
	id, err := rows.LastInsertId()
	if err != nil {
		return nil, err
	}

	room, err := GetRoom(id)
	if err != nil {
		return nil, err
	}

	return room, nil
}
