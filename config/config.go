package config

import (
	"fmt"

	"github.com/spf13/viper"
)

var v *viper.Viper

func init() {
	v = viper.New()
}

func InitConfig() {
	ReadFile()
	Connection(v.GetString("database.dbname"))
}

func ReadFile() {
	v.SetConfigName("config")
	v.SetConfigType("yaml")
	v.AddConfigPath(".")
	err := v.ReadInConfig()
	if err != nil {       
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
}

func GetConfig() *viper.Viper {
	return v
}

