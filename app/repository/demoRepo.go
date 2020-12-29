package repository

import (
	"github.com/gopher-lego/skeleton/config"
	"github.com/spf13/viper"
	"os"
)

func init()  {
	pwd, _ := os.Getwd()
	config.InitConf(pwd + "/../../setting")
}

func DemoRepo() string {
	return "DemoRepo::" + viper.GetString("name")
}
