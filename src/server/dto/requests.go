package dto

type CreateUserRequest struct {
	Username string `json:"username"`
}

type GetUserRequest struct {
	UserId string `json:"id"`
}

type UpdateUserRequest struct {
	UserId   int64  `json:"id"`
	Username string `json:"username"`
}

type CreateRoomRequest struct {
	User1 string `json:"user1"`
}

type UpdateRoomRequest struct {
	User2  string `json:"user2"`
	Winner string `json:"winner"`
}

type DeleteRoomRequest struct {
	RoomId int64 `json:"id"`
}

type UpdateRoomRequestMove struct {
	UserId int64 `json:"userid"`
	Row    int32 `json:"row"`
	Col    int32 `json:"col"`
}
