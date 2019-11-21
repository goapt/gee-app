package config

import (
	"github.com/ilibs/gosql/v2"
)

func Bootstrap() {
	//db connect
	_ = gosql.Connect(App.DB)
}
