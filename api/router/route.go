package router

import (
	"context"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/goapt/gee"
	"github.com/google/wire"

	"app/api/handler"
	"app/api/middleware"
)

func Engine(handler *handler.Handler, middleware *middleware.Middleware) http.Handler {
	router := gee.Default()
	// log middleware use for all handle
	router.POST("/login", handler.User.Login)

	// session middleware use for authorized handle
	api := router.Group("/api")
	api.Use(middleware.Session())
	{
		api.POST("/user", handler.User.Get)
	}

	// debug handler
	gee.DebugRoute(router.Engine)
	return router
}

func Setup(router http.Handler, addr string) {
	srv := &http.Server{
		Addr:    addr,
		Handler: router,
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

var ProviderSet = wire.NewSet(Engine)
