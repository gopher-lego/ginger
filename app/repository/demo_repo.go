package repository

import (
	"github.com/spf13/viper"
)

func Demo() string {
	return "Demo::" + viper.GetString("name")
}
