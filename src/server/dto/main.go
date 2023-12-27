package dto

import "time"

type Move struct {
	Userid int64 `json:"userid"`
	Row    int32 `json:"row"`
	Col    int32 `json:"col"`
}

type Room struct {
	RoomId      string    `json:"id"`
	CreatedDate time.Time `json:"date"`
	User1       string    `json:"user1"`
	User2       string    `json:"user2"`
	Winner      string    `json:"winner"`
	Moves       []*Move   `json:"moves"`
}

type User struct {
	UserId      string    `json:"id"`
	Username    string    `json:"username"`
	CreatedDate time.Time `json:"createdDate"`
}

func (u *User) String() string {
	return u.Username + ", with id: " + u.UserId
}
