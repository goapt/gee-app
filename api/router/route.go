package router

import (
	"context"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/goapt/gee"

	"app/api/handler"
	"app/api/middleware"
)

func Engine() *gee.Engine {
	router := gee.Default()
	// log middleware use for all handle
	router.POST("/login", handler.LoginHandle)

	// session middleware use for authorized handle
	api := router.Group("/api")
	api.Use(middleware.SessionMiddleware())
	{
		api.POST("/user", handler.UserHandle)
	}

	//debug handler
	gee.DebugRoute(router.Engine)
	return router
}

func Setup(addr string) {
	srv := &http.Server{
		Addr:    addr,
		Handler: Engine(),
	}
	log.Println("[HTTP] Server listen:" + addr)
	gee.RegisterShutDown(func(sig os.Signal) {
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
