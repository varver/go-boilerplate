package conf

import (
	"github.com/BurntSushi/toml"
	"github.com/varver/go-boilerplate/utils/logger"
	"log"
	"os"
)

type config struct {
	EnvMode    string
	GoMaxProcs int

	ServerPort       string
	ServerAppTagline string
	ServerAppName    string
	ServerAppVersion string

	DbDriver        string
	DbSource        string
	DbSourceDevMode string

	TemplatePath string
	DomainName   string
	UploadDir    string
	CookieSecret string
	CookieName   string
}

var Setting config

const (
	DEV = "dev"
)

func init() {
	logger.Info("configurations init called off")
	dir, _ := os.Getwd()

	// order in which to search for config file
	files := []string{
		dir + "/dev.ini",
		dir + "/config.ini",
		dir + "/conf/dev.ini",
		dir + "/conf/config.ini",
	}

	for _, f := range files {
		if _, err := toml.DecodeFile(f, &Setting); err == nil {
			logger.Info("Loaded configuration %s", f)
			break
		}
	}

	if len(Setting.ServerPort) < 3 {
		logger.Emerg("Problem in loading configuration file.")
		log.Panicln("Configuration files are not loaded properly, problem in finding port to run application.")
	}

	// make changes as per environment settings
	if Setting.EnvMode == DEV {
		// use dev mode database settings
		Setting.DbSource = Setting.DbSourceDevMode
	}

}

func GetConfig() *config {
	return &Setting
}
