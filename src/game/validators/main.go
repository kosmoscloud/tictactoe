package validators

import (
	"log"
	db "tictactoe-service/database"
	"tictactoe-service/server/dto"
)

func IsMoveValid(roomId int64, givenMove *dto.Move) bool {
	moves, err := db.GetMoves(roomId)
	if err != nil {
		log.Default().Println(err)
		return false
	}

	for _, move := range moves {
		if move.Row == givenMove.Row && move.Col == givenMove.Col {
			return false
		}
	}

	return true
}
