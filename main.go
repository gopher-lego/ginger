package main

import (
	"bytes"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"

	"github.com/gopher-lego/skeleton/route"

	"github.com/gopher-lego/skeleton/config"
)

var AppConf = viper.New()

/**
 * curl -X GET -d "s=a"  http://localhost:8090/api/ping
 */
func main() {

	if gin.Mode() == gin.ReleaseMode {

		// https://github.com/go-bindata/go-bindata#accessing-an-asset
		filename := "setting/app." + gin.Mode() + ".json"
		bytesContent, err := Asset(filename)
		if err != nil {
			panic("Asset() can not found setting file")
		}
		// https://github.com/spf13/viper#reading-config-from-ioreader
		AppConf.SetConfigType("json")
		if err := AppConf.ReadConfig(bytes.NewBuffer(bytesContent)); err != nil {
			panic("Something error:" + err.Error())
		}

	} else {

		filenameWithoutExt := "app." + gin.Mode()
		AppConf.SetConfigName(filenameWithoutExt)
		AppConf.SetConfigType("json")
		AppConf.AddConfigPath("./setting")
		if err := AppConf.ReadInConfig(); err != nil {
			panic("Using config file:" + AppConf.ConfigFileUsed())
		}

	}

	// Assign config for outside using
	config.Set(AppConf)

	// Framework engine
	engine := gin.Default()

	// Routes
	route.Set(engine)

	// Run
	engine.Run(AppConf.Get("server.port").(string)) // listen and serve on 0.0.0.0:port
}
