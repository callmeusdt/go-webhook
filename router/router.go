package router

import (
	"github.com/gin-gonic/gin"
	"hook/app"
	"hook/app/default"
)

func SetupRouter() *gin.Engine {
	app.GIN.POST("/hook", _default.Hook)
	app.GIN.GET("/ping", _default.Pong)

	return app.GIN
}
