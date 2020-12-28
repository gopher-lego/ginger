package repository

import (
	"fmt"
	"github.com/gopher-lego/skeleton/config"
	"github.com/spf13/viper"
	"os"
)

func init()  {
	pwd, _ := os.Getwd()
	config.InitConf(pwd + "/../../setting")
}

func DemoRepo() {
	fmt.Println(viper.Get("name"), "::DemoRepo")
}
