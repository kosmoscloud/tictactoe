package errors

import "tictactoe-service/server/dto"

type MoveAlreadyExists struct {
	AlreadyExistingMove *dto.Move
}

func (e *MoveAlreadyExists) Error() string {
	return "cannot add move - move already exists: " + e.AlreadyExistingMove.String()
}

type UserAlreadyExistsInRoom struct{}

func (e *UserAlreadyExistsInRoom) Error() string {
	return "cannot add user to room - user already exists in room"
}
