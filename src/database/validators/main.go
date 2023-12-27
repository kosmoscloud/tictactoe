package validators

import (
	"log"
	"tictactoe-service/server/dto"
)

func IsMoveValid(roomId int64, givenMove *dto.Move) bool {
	if givenMove.Row < 0 || givenMove.Row > 2 {
		log.Default().Println("Invalid row: ", givenMove.Row)
		return false
	}
	if givenMove.Col < 0 || givenMove.Col > 2 {
		log.Default().Println("Invalid col: ", givenMove.Col)
		return false
	}
	return true
}
