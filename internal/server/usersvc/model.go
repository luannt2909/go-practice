package usersvc

import "go-practice/pkg/user"

type CreateUserRequest struct {
	Name   string      `json:"name"`
	Gender user.Gender `json:"gender"`
}

type CreateUserResponse struct {
	Code    int64     `json:"code"`
	Message string    `json:"message"`
	Data    user.User `json:"data"`
}

type GetUserByIDResponse struct {
	Code    int64     `json:"code"`
	Message string    `json:"message"`
	Data    user.User `json:"data"`
}
