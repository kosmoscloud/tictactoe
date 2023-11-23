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
