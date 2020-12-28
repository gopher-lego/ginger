package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gopher-lego/skeleton/route"
	"github.com/spf13/viper"

	"github.com/gopher-lego/skeleton/config"
)

/**
 * curl -X GET -d "s=a"  http://localhost:8090/api/ping
 */
func main() {

	//if gin.Mode() == gin.ReleaseMode {
	//
	//	// https://github.com/go-bindata/go-bindata#accessing-an-asset
	//	filename := "setting/app." + gin.Mode() + ".json"
	//	bytesContent, err := Asset(filename)
	//	if err != nil {
	//		panic("Asset() can not found setting file")
	//	}
	//	// https://github.com/spf13/viper#reading-config-from-ioreader
	//	AppConf.SetConfigType("json")
	//	if err := AppConf.ReadConfig(bytes.NewBuffer(bytesContent)); err != nil {
	//		panic("Something error:" + err.Error())
	//	}
	//
	//}

	config.InitConf("./setting")

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
