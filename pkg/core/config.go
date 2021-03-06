package core

import (
	"fmt"

	"github.com/spf13/viper"
)

type CFG_Db struct {
	Dsn string
}

type CFG_Log struct {
	Level string
}

type CFG_Notification struct {
	Smtp *CFG_Smtp
}

type CFG_Smtp struct {
	Host     string
	Port     int
	Username string
	Password string
}

type CFG_Server struct {
	Port int
}

type CFG_Site struct {
	Domain        string
	DefaultAvatar string
	Brand         string
	Footer        string
	Name          string
}

type cfg struct {
	Db           *CFG_Db
	Log          *CFG_Log
	Notification *CFG_Notification
	Server       *CFG_Server
	Site         *CFG_Site
}

var CFG *cfg

func LoadConfig(args []string) {
	configPath := "config.yaml"
	if len(args) >= 2 {
		configPath = args[1]
	}
	viper.SetDefault("db.dsn", "root:123456@tcp(127.0.0.1:3306)/gorum?charset=utf8mb4&parseTime=True&loc=Local")
	viper.SetDefault("server.port", 8080)
	viper.SetDefault("log.level", "info")
	viper.SetDefault("site.defaultAvatar", "/assets/avatar.svg")
	viper.SetDefault("site.name", "未命名站点")
	viper.SetDefault("site.brand", viper.GetString("site.name"))
	viper.SetDefault("notification", make(map[string]string))
	viper.SetConfigType("yaml")
	viper.SetConfigFile(configPath)
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
	c := cfg{}
	err = viper.Unmarshal(&c)
	if err != nil {
		panic(err)
	}
	CFG = &c
}

func (cfg *cfg) String(key string) string {
	return viper.GetString(key)
}

func (cfg *cfg) Int(key string) int {
	return viper.GetInt(key)
}
