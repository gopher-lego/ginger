package config

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

// Initial in main.go to bind setting in binary file with different method !

func InitConf() {
	// File name without extension
	filenameWithoutExt := "app." + gin.Mode()

	viper.SetConfigName(filenameWithoutExt)
	viper.SetConfigType("json")
	viper.AddConfigPath("./setting") // Project dir
	if err := viper.ReadInConfig(); err != nil {
		panic("Using config file:" + viper.ConfigFileUsed())
	}
}
