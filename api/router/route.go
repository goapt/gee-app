package router

import (
	"context"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/goapt/gee"
	"github.com/goapt/golib/osutil"

	"app/api/handler"
	"app/api/middleware"
)

func Engine() *gin.Engine {
	binding.Validator = new(gee.DefaultValidator)

	router := gin.Default()
	// log middleware use for all handle
	router.POST("/login", gee.Handle(handler.LoginHandle))

	// session middleware use for authorized handle
	api := router.Group("/api")
	api.Use(middleware.TokenMiddleware())
	{
		api.POST("/user", gee.Handle(handler.UserHandle))
	}

	//debug handler
	gee.DebugRoute(router)
	return router
}

func Setup(addr string) {
	srv := &http.Server{
		Addr:    addr,
		Handler: Engine(),
	}

	osutil.RegisterShutDown(func(sig os.Signal) {
		ctxw, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		_ = srv.Close()
		if err := srv.Shutdown(ctxw); err != nil {
			log.Fatal("HTTP Server Shutdown:", err)
		}
		log.Println("HTTP Server exiting")
	})

	// service connections
	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("HTTP listen: %s\n", err)
	}
}
