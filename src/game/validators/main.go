package validators

import (
	"log"
	db "tictactoe-service/database"
	"tictactoe-service/game/errors"
	"tictactoe-service/server/dto"
)

func IsMoveLegal(roomId int64, givenMove *dto.Move) (bool, error) {
	moves, err := db.GetMoves(roomId)
	if err != nil {
		log.Default().Println(err)
		return false, &(errors.InvalidMoveValue{InvalidMove: givenMove})
	}

	for _, move := range moves {
		if move.Row == givenMove.Row && move.Col == givenMove.Col {
			return false, &(errors.MoveAlreadyExists{AlreadyExistingMove: move})
		}
	}

	if moves[len(moves)-1].UserId == givenMove.UserId {
		return false, &(errors.UserJustMoved{UserId: givenMove.UserId})
	}

	return true, nil
}

func IsMoveWinning(roomId int64, givenMove *dto.Move) bool {
	moves, err := db.GetMoves(roomId)
	if err != nil {
		log.Default().Println(err)
		return false
	}

	// check row
	rowCount := 0
	for _, move := range moves {
		if move.Row == givenMove.Row && move.UserId == givenMove.UserId {
			rowCount++
		}
	}

	// check col
	colCount := 0
	for _, move := range moves {
		if move.Col == givenMove.Col && move.UserId == givenMove.UserId {
			colCount++
		}
	}

	// check diagonal
	diagonalCount := 0
	for _, move := range moves {
		if move.Row == move.Col && move.UserId == givenMove.UserId {
			diagonalCount++
		}
	}

	// check anti-diagonal
	antiDiagonalCount := 0
	for _, move := range moves {
		if move.Row+move.Col == 2 && move.UserId == givenMove.UserId {
			antiDiagonalCount++
		}
	}

	return rowCount == 3 || colCount == 3 || diagonalCount == 3 || antiDiagonalCount == 3
}
