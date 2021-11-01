package config

import (
	"os"
	"sync"

	"github.com/labstack/gommon/log"
	"github.com/spf13/viper"
)

type AppConfig struct {
	AppPort        int    `mapstructure:"app_port"`
	AppHost        string `mapstructure:"app_host"`
	AppEnvironment string `mapstructure:"app_environment"`
	DbDriver       string `mapstructure:"db_driver"`
	DbHost         string `mapstructure:"db_host"`
	DbPort         int    `mapstructure:"db_port"`
	DbUsername     string `mapstructure:"db_username"`
	DbPassword     string `mapstructure:"db_password"`
	DbName         string `mapstructure:"db_name"`
	JWTConfig      string `mapstructure:"jwt_secret"`
}

var lock = &sync.Mutex{}
var appConfig *AppConfig

func GetConfig() *AppConfig {
	if appConfig != nil {
		return appConfig
	}
	lock.Lock()
	defer lock.Unlock()

	if appConfig != nil {
		return appConfig
	}
	appConfig = InitConfig()

	return appConfig
}

func InitConfig() *AppConfig {
	var defaultConfig AppConfig

	defaultConfig.AppPort = 9092
	defaultConfig.AppHost = "localhost"
	defaultConfig.AppEnvironment = ""
	defaultConfig.DbDriver = "mongodb"
	defaultConfig.DbHost = "localhost"
	defaultConfig.DbPort = 3306
	defaultConfig.DbUsername = "root"
	defaultConfig.DbPassword = "1234"
	defaultConfig.DbName = "altaecom"
	defaultConfig.JWTConfig = "altaecom"

	var (
		err     error
		curPath string
	)

	curPath, err = os.Getwd()
	if err != nil {
		log.Info("failed get current directory")
		return &defaultConfig
	}
	// viper.SetConfigFile(curPath + "/config/.env")
	viper.SetConfigFile(curPath + "/../config/.env")
	err = viper.ReadInConfig()
	if err != nil {
		log.Info("failed read env file")
		return &defaultConfig
	}

	var finalConfig AppConfig
	err = viper.Unmarshal(&finalConfig)
	if err != nil {
		log.Info("failed bind config")
		return &defaultConfig
	}

	return &finalConfig
}
