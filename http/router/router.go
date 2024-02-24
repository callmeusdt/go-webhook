package router

import (
	"github.com/gin-gonic/gin"
	"go-webhook/http"
	"go-webhook/http/controller"
	"go-webhook/http/middleware"
)

func SetupRouter() *gin.Engine {
	_group := http.SERVER.Group("/")
	_group.Use(middleware.LogRequest(), middleware.LogResponse())

	_group.POST("hook", controller.Hook)
	_group.GET("ping", controller.Ping)

	return http.SERVER
}
