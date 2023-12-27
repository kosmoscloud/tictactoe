package database

import (
	"tictactoe-service/server/dto"
)

func CreateMove(roomId int64, givenMove *dto.Move) (*dto.Move, error) {
	rows, err := DB.Exec("INSERT INTO moves (room_id, user_id, row_, col_) VALUES (?, ?, ?, ?)",
		roomId, givenMove.UserId, givenMove.Row, givenMove.Col)
	if err != nil {
		return nil, err
	}
	id, err := rows.LastInsertId()
	if err != nil {
		return nil, err
	}

	move, err := GetMove(id)
	if err != nil {
		return nil, err
	}

	return move, nil

}

func GetMove(id int64) (*dto.Move, error) {

	row := DB.QueryRow("SELECT id, room_id, user_id, row_, col_ FROM moves WHERE id=?", id)
	move := &dto.Move{}
	var idmove, room_id int64
	err := row.Scan(&idmove, &room_id, &move.UserId, &move.Row, &move.Col)

	if err != nil {
		return nil, err
	}

	return move, nil
}

func GetMoves(roomId int64) ([]*dto.Move, error) {
	moves := []*dto.Move{}
	rows, err := DB.Query("SELECT user_id, row_, col_ FROM moves WHERE room_id=?", roomId)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		move := &dto.Move{}
		err := rows.Scan(&move.UserId, &move.Row, &move.Col)
		if err != nil {
			return nil, err
		}
		moves = append(moves, move)
	}
	return moves, nil
}
