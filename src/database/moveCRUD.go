package database

import (
	"tictactoe-service/server/dto"
)

func createdMove(roomId int64, userId int64, row int64, col int64) (*dto.Move, *dto.Room, error) {
	rows, err := DB.Exec("INSERT INTO moves (room_id, user_id, row_, col) VALUES (?, ?, ?, ?)", roomId, userId, row, col)
	if err != nil {
		return nil, nil, err
	}
	id, err := rows.LastInsertId()
	if err != nil {
		return nil, nil, err
	}

	move, room, err := GetMove(id)
	if err != nil {
		return nil, nil, err
	}

	return move, room, nil

}

func GetMove(id int64) (*dto.Move, *dto.Room, error) {
	move := &dto.Move{}
	room := &dto.Room{}
	row := DB.QueryRow("SELECT id, room_id, user_id, row_, col FROM moves WHERE id=?", id)
	err := row.Scan(&room.RoomId, &room.User1, &move.Row, &move.Col)
	if err != nil {
		return nil, nil, err
	}
	return move, room, nil
}

func GetMoves(roomId int64) ([]*dto.Move, *dto.Room, error) {
	moves := []*dto.Move{}
	room := &dto.Room{}
	rows, err := DB.Query("SELECT id, room_id, user_id, row_, col FROM moves WHERE room_id=?", roomId)
	if err != nil {
		return nil, nil, err
	}
	for rows.Next() {
		move := &dto.Move{}
		err := rows.Scan(&move.Row, &move.Col)
		if err != nil {
			return nil, nil, err
		}
		moves = append(moves, move)
	}
	return moves, room, nil
}
