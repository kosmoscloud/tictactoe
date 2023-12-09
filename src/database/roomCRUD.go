package database

import (
	"tictactoe-service/server/dto"
	"time"
)

func GetRoom(id int64) (*dto.Room, error) {
	room := &dto.Room{}
	row := DB.QueryRow("SELECT id, created, user1, user2, winner FROM rooms WHERE id=?", id)
	err := row.Scan(&room.RoomId, &room.CreatedDate, &room.User1, &room.User2, &room.Winner)
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

func UpdateRoom(RoomId int64, winner int64) (*dto.Room, error) {
	_, err := DB.Exec("UPDATE rooms SET winner=? WHERE id=?", winner, RoomId)
	if err != nil {
		return nil, err
	}
	room, err := GetRoom(RoomId)
	if err != nil {
		return nil, err
	}

	return room, nil

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
	//_, err = DB.Exec("DELETE FROM moves WHERE room_id=?", id)

	return room, nil
}
