package dto

type CreateUserRequest struct {
	Username string `json:"username"`
}

type GetUserRequest struct {
	UserId string `json:"id"`
}
