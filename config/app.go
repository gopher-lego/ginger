package config

import (
	"github.com/spf13/viper"
)

var AppConf *viper.Viper

// Called in main.go, then config.AppConf.Get("xx") used in other go files
func Set(v *viper.Viper) {
	AppConf = v
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
