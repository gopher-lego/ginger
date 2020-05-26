package main

import (
	"bytes"
	"github.com/gin-gonic/gin"
	// "github.com/gopher-lego/skeleton/app/model"
	"github.com/gopher-lego/skeleton/config"
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
)

var AppConfig = viper.New()

func main() {

	if gin.Mode() == gin.ReleaseMode {

		// https://github.com/go-bindata/go-bindata#accessing-an-asset
		filename := "setting/app." + gin.Mode() + ".json"
		bytesContent, err := Asset(filename)
		if err != nil {
			panic("Asset() can not found setting file")
		}
		// https://github.com/spf13/viper#reading-config-from-ioreader
		AppConfig.SetConfigType("json")
		if err := AppConfig.ReadConfig(bytes.NewBuffer(bytesContent)); err != nil {
			panic("Something error:" + err.Error())
		}

	} else {

		filenameWithoutExt := "app." + gin.Mode()
		AppConfig.SetConfigName(filenameWithoutExt)
		AppConfig.SetConfigType("json")
		AppConfig.AddConfigPath("./setting")
		if err := AppConfig.ReadInConfig(); err != nil {
			panic("Using config file:" + AppConfig.ConfigFileUsed())
		}

	}

	// Assign config for outside using
	config.Set(AppConfig)

	// ================================

	db := config.MySqlInit()

	db.LogMode(true)

	panic("Please annotate code below")

	// Migrate
	// db.AutoMigrate(&model.User{})
	defer db.Close()
}
