package config

import (
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/goapt/gee"
	"github.com/ilibs/gosql/v2"
	"github.com/pelletier/go-toml"
)

type app struct {
	Env         string                   `toml:"env"`
	AppName     string                   `toml:"app_name"`
	StoragePath string                   `toml:"storage_path"`
	Debug       string                   `toml:"debug"`
	TokenSecret string                   `toml:"token_secret"`
	DB          map[string]*gosql.Config `toml:"database"`
	StartTime   time.Time
	IsTesting   bool
}

var App = &app{
	StartTime: time.Now(),
}

func init() {
	gee.ArgsInit()
	load(gee.ExtCliArgs)
}

func load(args map[string]string) {
	appPath := args["config"]

	if appPath == "" {
		//由于go test执行路径是临时目录，因此寻找配置文件要从编译路径查找
		if strings.HasSuffix(os.Args[0], ".test") {
			App.IsTesting = true
			_, file, _, _ := runtime.Caller(0)
			appPath = filepath.Dir(filepath.Dir(file))
		} else {
			appPath = "./"
		}
	}

	conf, err := toml.LoadFile(filepath.Join(appPath, "config.toml"))
	if err != nil {
		log.Fatalf("config error %s", err.Error())
	}

	if err := conf.Unmarshal(App); err != nil {
		log.Fatal("config unmarshal error:", err)
	}

	if !filepath.IsAbs(App.StoragePath) {
		App.StoragePath = filepath.Join(appPath, App.StoragePath)
	}

	if App.Env == "local" {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	//debug model
	if args["debug"] != "" {
		App.Debug = args["debug"]
	}

	if args["show-sql"] == "on" {
		for _, d := range App.DB {
			d.ShowSql = true
		}
	} else if args["show-sql"] == "off" {
		for _, d := range App.DB {
			d.ShowSql = false
		}
	}
}
