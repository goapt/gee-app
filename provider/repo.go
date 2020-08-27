package provider

import (
	"github.com/google/wire"

	"app/provider/user/repo"
)

var RepoSet = wire.NewSet(repo.NewUsers)
