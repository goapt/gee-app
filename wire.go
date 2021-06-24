// +build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/urfave/cli/v2"

	"app/api/handler"
	"app/api/middleware"
	"app/api/router"
	"app/cmd"
	"app/connect"
	"app/provider"
)

var providerSet = wire.NewSet(
	connect.ProviderSet,
	cmd.ProviderSet,
	router.ProviderSet,
	middleware.ProviderSet,
	handler.ProviderSet,
	provider.RepoSet,
)

func Initialize() cli.Commands {
	panic(wire.Build(providerSet))
}
