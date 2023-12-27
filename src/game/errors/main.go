package errors

import (
	"strconv"
	"tictactoe-service/server/dto"
)

type MoveAlreadyExists struct {
	AlreadyExistingMove *dto.Move
}

func (e *MoveAlreadyExists) Error() string {
	return "cannot add move - move already exists: " + e.AlreadyExistingMove.String()
}

type InvalidMoveValue struct {
	InvalidMove *dto.Move
}

func (e *InvalidMoveValue) Error() string {
	return "cannot add move - invalid move: " + e.InvalidMove.String() + " - row and col must be between 0 and 2"
}

type UserJustMoved struct {
	UserId int64
}

func (e *UserJustMoved) Error() string {
	return "cannot add move - user with id: " + strconv.FormatInt(e.UserId, 10) + " just moved"
}
