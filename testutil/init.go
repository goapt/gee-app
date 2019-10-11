package testutil

import (
	"app/config"

	_ "github.com/go-sql-driver/mysql"
	"github.com/ilibs/gosql/v2"
)

func init() {
	//db connect
	_ = gosql.Connect(config.App.DB)
}
