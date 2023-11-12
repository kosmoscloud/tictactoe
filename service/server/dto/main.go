package dto

import "google.golang.org/genproto/googleapis/type/datetime"

type Move struct {
	Row int32 `json:"row"`
	Col int32 `json:"col"`
}

type Room struct {
	RoomId      string            `json:"id"`
	CreatedDate datetime.DateTime `json:"date"`
	User1       string            `json:"user1"`
	User2       string            `json:"user2"`
	Winner      string            `json:"winner"`
	Moves       []Move            `json:"moves"`
}

type User struct {
	UserId      string            `json:"id"`
	Username    string            `json:"username"`
	CreatedDate datetime.DateTime `json:"createdDate"`
}
