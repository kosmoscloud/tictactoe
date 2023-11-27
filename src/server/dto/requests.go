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
	User2 string `json:"user2"`
}
