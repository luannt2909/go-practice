package main

import (
	"context"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/fx"

	"go-practice/internal/di/cachefx"
	"go-practice/internal/di/configfx"
	"go-practice/internal/di/dbfx"
	"go-practice/internal/di/ginfx"
	"go-practice/internal/di/redisfx"
	"go-practice/internal/di/serverfx"
	"go-practice/internal/di/userfx"
	"go-practice/internal/server/usersvc"
)

func main() {
	app := fx.New(
		configfx.Initialize("config.yaml", "configs"),
		ginfx.Module,
		dbfx.Module,
		redisfx.Module,
		cachefx.Module,
		userfx.Module,
		serverfx.Module,
		fx.Invoke(
			registerService,
			startServer),
	)
	app.Run()
}

func registerService(ginEngine *gin.Engine, userSvcRouter usersvc.Router) {
	gGroup := ginEngine.Group("api/v1")
	userSvcRouter.Register(gGroup)
}

func startServer(ginEngine *gin.Engine, lifecycle fx.Lifecycle) {
	server := http.Server{
		Addr:    ":3000",
		Handler: ginEngine,
	}
	lifecycle.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			err := server.ListenAndServe()
			if err != nil && !errors.Is(err, http.ErrServerClosed) {
				return err
			}
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return server.Shutdown(ctx)
		},
	})
}
