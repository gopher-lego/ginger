package main

import (
	_ "embed"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/gopher-lego/ginger/route"
	"github.com/spf13/viper"

	"github.com/gopher-lego/ginger/config"
)

/**
 * $ go run main.go bindata.go
 * $ curl -X GET -d "s=a"  http://localhost:8090/api/ping
 */
func main() {
	// Load configure file
	config.InitConfAsset(mainPath, bytesContent)

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

var mainPath = func() string {
	pwd, _ := os.Getwd()
	return pwd
}()

//go:embed setting/app.release.json
var bytesContent []byte