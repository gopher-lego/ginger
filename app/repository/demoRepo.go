package repository

import (
	"github.com/spf13/viper"
)

func DemoRepo() string {
	return "DemoRepo::" + viper.GetString("name")
}
