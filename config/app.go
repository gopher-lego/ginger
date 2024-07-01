package config

import (
	"bytes"
	_ "embed"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

// Initial in main.go to bind setting into bytes !

// Then use example: viper.Get("cache.freecache.size")
func InitConfAsset(mainPath string, bytesContent []byte) {
	if gin.Mode() == gin.ReleaseMode {
		InitReleaseConf(bytesContent)
	} else {
		InitOtherConf(mainPath + "/setting")
	}
}

func InitReleaseConf(bytesContent []byte) {
	// https://github.com/spf13/viper#reading-config-from-ioreader
	viper.SetConfigType("json")
	if err := viper.ReadConfig(bytes.NewBuffer(bytesContent)); err != nil {
		panic("viper.ReadConfig something error:" + err.Error())
	}
}

// Path can be relative or absolute.
func InitOtherConf(settingPath string) {
	// File name without extension
	filenameWithoutExt := "app." + gin.Mode()

	// viper.New usage is for multi instance case
	viper.SetConfigName(filenameWithoutExt)
	viper.SetConfigType("json")
	viper.AddConfigPath(settingPath)
	if err := viper.ReadInConfig(); err != nil {
		panic("Using config file:" + viper.ConfigFileUsed())
	}
}