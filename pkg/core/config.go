package core

import (
	"fmt"

	"github.com/spf13/viper"
)

type cfg struct {
}

var CFG *cfg

func LoadConfig() {
	viper.SetDefault("db.dsn", "root:123456@tcp(127.0.0.1:3306)/gorum?charset=utf8mb4&parseTime=True&loc=Local")
	viper.SetConfigName("config") // name of config file (without extension)
	viper.SetConfigType("yaml")   // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath(".")      // optionally look for config in the working directory
	err := viper.ReadInConfig()   // Find and read the config file
	if err != nil {               // Handle errors reading the config file
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
	CFG = &cfg{}
}

func (cfg *cfg) String(key string) string {
	return viper.GetString(key)
}
