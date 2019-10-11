package router

import (
	"github.com/gin-gonic/gin"
	"github.com/goapt/gee"

	"app/handler"
	"app/router/middleware"
)

func Route(router *gin.Engine) {

	router.POST("/login", gee.Handle(handler.LoginHandle))

	// session middleware use for authorized handle
	api := router.Group("/api")
	api.Use(middleware.TokenMiddleware())
	{
		api.POST("/user", gee.Handle(handler.UserHandle))
	}

	//debug handler
	gee.DebugRoute(router)
}
