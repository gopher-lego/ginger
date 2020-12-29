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
	pwd, _ := os.Getwd()
	settingPath := pwd + "/setting"

	// Load configure file
	if gin.Mode() == gin.ReleaseMode {
		InitConfAsset(settingPath)
	} else {
		// Debug or Test mode
		config.InitConf(settingPath)
	}

	// Memory cache
	config.NewFreeCache()

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

func InitConfAsset(confPath string) {
	// https://github.com/go-bindata/go-bindata#accessing-an-asset
	filename := "app." + gin.Mode() + ".json"
	bytesContent, err := Asset(confPath + "/" + filename)
	if err != nil {
		panic("Asset() can not found setting file")
	}
	// https://github.com/spf13/viper#reading-config-from-ioreader
	viper.SetConfigType("json")
	if err := viper.ReadConfig(bytes.NewBuffer(bytesContent)); err != nil {
		panic("viper.ReadConfig something error:" + err.Error())
	}
}
