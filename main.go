package main

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"github.com/gopher-lego/ginger/route"
	"github.com/spf13/viper"
	"os"

	"github.com/gopher-lego/ginger/config"
)

/**
 * $ go run main.go bindata.go
 * $ curl -X GET -d "s=a"  http://localhost:8090/api/ping
 */
func main() {
	// Load configure file
	if gin.Mode() == gin.ReleaseMode {
		InitConfAsset()
	} else {
		// Debug or Test mode
		pwd, _ := os.Getwd()
		settingPath := pwd + "/setting"
		config.InitConf(settingPath)
	}

	// config.MySqlInit()

	// Memory cache
	config.NewFreeCache()

	// simple single-node memory cache
	config.NewGoCache()

	// Framework engine
	engine := gin.Default()

	// Routes
	route.Set(engine)

	// Run
	err := engine.Run(viper.GetString("server.port")) // listen and serve on 0.0.0.0:port
	if err != nil {
		panic(err.Error())
	}
}

// File path will not change
func InitConfAsset() {
	// https://github.com/go-bindata/go-bindata#accessing-an-asset
	filename := "setting/app." + gin.Mode() + ".json"
	bytesContent, err := Asset(filename)
	if err != nil {
		panic("Asset() can not found setting file")
	}
	// https://github.com/spf13/viper#reading-config-from-ioreader
	viper.SetConfigType("json")
	if err := viper.ReadConfig(bytes.NewBuffer(bytesContent)); err != nil {
		panic("viper.ReadConfig something error:" + err.Error())
	}
}
