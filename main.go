package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"go.uber.org/fx"

	"go-practice/internal/di/configfx"
	"go-practice/internal/di/ginfx"
	"go-practice/internal/di/serverfx"
	"go-practice/internal/server/usersvc"
)

func main() {
	app := fx.New(
		configfx.Initialize("config.yaml", "configs"),
		ginfx.Module,
		// dbfx.Module,
		// redisfx.Module,
		// cachefx.Module,
		// userfx.Module,
		serverfx.Module,
		fx.Invoke(
			// registerService,
			startServer),
	)
	app.Run()
}

func registerService(ginEngine *gin.Engine, userSvcRouter usersvc.Router) {
	gGroup := ginEngine.Group("api/v1")
	userSvcRouter.Register(gGroup)
}

func startServer(ginEngine *gin.Engine, lifecycle fx.Lifecycle) {
	port := viper.GetString("PORT")
	server := http.Server{
		Addr:    ":" + port,
		Handler: ginEngine,
	}
	ginEngine.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})
	lifecycle.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			fmt.Println("run on port:", port)
			go func() {
				if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
					fmt.Errorf("failed to listen and serve from server: %v", err)
				}
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return server.Shutdown(ctx)
		},
	})
}
