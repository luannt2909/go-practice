package usersvc

import (
	"context"

	"go-practice/pkg/user"
)

type Controller interface {
	CreateUser(c context.Context, request CreateUserRequest) (user.User, error)
	GetUserByID(c context.Context, id uint64) (user.User, error)
}

type controller struct {
	userRepo user.Repository
}

func (ctrl *controller) CreateUser(c context.Context, request CreateUserRequest) (user.User, error) {
	user := user.User{
		Name:   request.Name,
		Gender: request.Gender,
	}
	err := ctrl.userRepo.Create(c, &user)
	return user, err
}

func (ctrl *controller) GetUserByID(c context.Context, id uint64) (user.User, error) {
	return ctrl.userRepo.GetByID(c, id)
}

func NewController(userRepo user.Repository) Controller {
	return &controller{userRepo: userRepo}
}
