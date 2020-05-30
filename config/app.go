package config

import (
	"github.com/spf13/viper"
)

var AppConf *viper.Viper

// Called in main.go, then config.AppConf.Get("xx") used in other go files
func Global(v *viper.Viper) {
	AppConf = v
}

// For easy use outside by config.Get('xx')
func Get(key string) string {
	return AppConf.Get(key).(string)
}

// Initial in main.go to bind setting in binary file !

// func init() {
// 	// File name without extension
// 	filenameWithoutExt := "app.debug"
//
// 	AppConf.SetConfigName(filenameWithoutExt)
// 	AppConf.SetConfigType("json")
// 	AppConf.AddConfigPath("./setting") // Project dir
// 	if err := AppConf.ReadInConfig(); err != nil {
// 		panic("Using config file:" + AppConf.ConfigFileUsed())
// 	}
// }
