package config

import (
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/goapt/envconf"
	"github.com/goapt/gee"
	"github.com/goapt/redis"
	"github.com/ilibs/gosql/v2"
)

type app struct {
	Env         string                   `toml:"env"`
	AppName     string                   `toml:"app_name"`
	StoragePath string                   `toml:"storage_path"`
	Debug       string                   `toml:"debug"`
	TokenSecret string                   `toml:"token_secret"`
	DB          map[string]*gosql.Config `toml:"database"`
	Redis       map[string]redis.Config  `toml:"redis"`
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

// 判断是否为测试执行
func isTestMode() bool {
	// test执行文件的路径后缀带.test，生产环境的可执行文件，不可能带.test后缀
	if strings.HasSuffix(os.Args[0], ".test") {
		return true
	}

	testVars := map[string]bool{
		"-test.v":   true,
		"-test.run": true,
	}

	for _, str := range os.Args {
		if testVars[str] {
			return true
		}
	}

	return false
}

func load(args map[string]string) {
	appPath := args["config"]

	if appPath == "" {
		if isTestMode() {
			App.IsTesting = true
			_, file, _, _ := runtime.Caller(0)
			appPath = filepath.Dir(filepath.Dir(file))
		} else {
			appPath = "./"
		}
	}

	conf, err := envconf.New(filepath.Join(appPath, "config.toml"))
	if err != nil {
		log.Fatalf("config error %s", err.Error())
	}

	// load env config
	if err := conf.Env(filepath.Join(appPath, ".env")); err != nil {
		log.Fatal("config env error:", err)
	}

	if err := conf.Unmarshal(App); err != nil {
		log.Fatal("config unmarshal error:", err)
	}

	if !filepath.IsAbs(App.StoragePath) {
		App.StoragePath = filepath.Join(appPath, App.StoragePath)
	}

	if App.Env == "local" && !isTestMode() {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	// debug model
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
