package database

import (
	"tictactoe-service/server/dto"
	"time"
)

func GetRoom(id int64) (*dto.Room, error) {
	room := &dto.Room{}
	row := DB.QueryRow("SELECT id, created, user1, user2 FROM rooms WHERE id=?", id)
	err := row.Scan(&room.RoomId, &room.CreatedDate, &room.User1, &room.User2)
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

func UpdateRoom(move *dto.Move) (*dto.Room, error) {
	return nil, nil
}

func DeleteRoom(id int64) (*dto.Room, error) {
	room, err := GetRoom(id)
	if err != nil {
		return nil, err
	}

	_, err = DB.Exec("DELETE FROM rooms WHERE id=?", id)
	if err != nil {
		return nil, err
	}

	return room, nil
}
