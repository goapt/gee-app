package cmd

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
	"github.com/urfave/cli"

	"app/router"
)

var HTTPCmd = cli.Command{
	Name:  "http",
	Usage: "http command eg: ./app http --addr=:8080",
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:  "addr",
			Usage: "http listen ip:port",
		},
	},
	Action: func(ctx *cli.Context) error {
		if !ctx.IsSet("addr") {
			_ = ctx.Set("addr", ":8080")
		}

		binding.Validator = new(defaultValidator)

		serv := gin.Default()
		srv := &http.Server{
			Addr:    ctx.String("addr"),
			Handler: serv,
		}

		//router
		router.Route(serv)

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

		return nil
	},
}

func init() {
	gee.RegisterCommand(HTTPCmd)
}
