package serverfx

import (
	"go-practice/internal/server/usersvc"
	"go-practice/pkg/user"
)

func provideUserController(userRepo user.Repository) usersvc.Controller {
	return usersvc.NewController(userRepo)
}

func provideUserRouter(userCtrl usersvc.Controller) usersvc.Router {
	return usersvc.NewRouter(userCtrl)
}
