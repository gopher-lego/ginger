package config

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

// Initial in main.go to bind setting in binary file with different method !

// Path can be relative or absolute.
func InitConf(confPath string) {
	// File name without extension
	filenameWithoutExt := "app." + gin.Mode()

	// viper.New usage is for multi instance case
	viper.SetConfigName(filenameWithoutExt)
	viper.SetConfigType("json")
	viper.AddConfigPath(confPath)
	if err := viper.ReadInConfig(); err != nil {
		panic("Using config file:" + viper.ConfigFileUsed())
	}
}