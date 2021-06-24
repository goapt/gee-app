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
	"github.com/goapt/logger"
	"github.com/goapt/redis"
	"github.com/ilibs/gosql/v2"
)

type Config struct {
	Env         string
	Path        string
	AppName     string                   `toml:"app_name"`
	StoragePath string                   `toml:"storage_path"`
	Debug       string                   `toml:"debug"`
	Log         logger.Config            `toml:"log"`
	DB          map[string]*gosql.Config `toml:"database"`
	Redis       map[string]redis.Config  `toml:"redis"`
	StartTime   time.Time
	IsTesting   bool
}

var App = &Config{
	StartTime: time.Now(),
}

func init() {
	gee.ArgsInit()
	load(gee.ExtCliArgs)
}

// isTestMode
func isTestMode() bool {
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

func getAppPath(path string) string {
	if path == "" {
		if isTestMode() {
			App.IsTesting = true
			_, file, _, _ := runtime.Caller(0)
			path = filepath.Dir(filepath.Dir(file))
		} else {
			path = "./"
		}
	}
	return path
}

func mustCheckError(err error) {
	if err != nil {
		log.Fatalf("config error %s", err.Error())
	}
}

func load(args map[string]string) {
	App.Path = getAppPath(args["config"])

	conf, err := envconf.New(filepath.Join(App.Path, "config.toml"))
	mustCheckError(err)

	if !App.IsTesting {
		err = conf.Env(filepath.Join(App.Path, ".env"))
		mustCheckError(err)
	}

	err = conf.Unmarshal(App)
	mustCheckError(err)

	if !filepath.IsAbs(App.StoragePath) {
		App.StoragePath = filepath.Join(App.Path, App.StoragePath)
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

	for _, d := range App.DB {
		d.ShowSql = args["show-sql"] == "on"
	}
}
