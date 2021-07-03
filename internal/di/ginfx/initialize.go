package ginfx

import (
	"github.com/gin-gonic/gin"
)

func provideGinEngine() *gin.Engine {
	return gin.Default()
}
