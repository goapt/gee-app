package cmd

import (
	"net/http"

	"github.com/urfave/cli"

	"app/api/router"
)

type HttpCmd cli.Command

func NewHttp(handler http.Handler) HttpCmd {
	return HttpCmd(cli.Command{
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
			router.Setup(handler, ctx.String("addr"))
			return nil
		},
	})
}
