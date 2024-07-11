package config

import (
	"fmt"
	"github.com/spf13/viper"
)

var (
	v    *viper.Viper
	lang string
)

func init() {
	v = viper.New()
	v.SetConfigName("config")
	v.SetConfigType("yaml")
	v.AddConfigPath(".")

	if err := v.ReadInConfig(); err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
}

func GetViper() *viper.Viper {
	return v
}

func SetLang(language string) {
	lang = language
}

func GetLang() string {
	return lang
}
