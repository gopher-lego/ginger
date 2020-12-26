package config

import (
	"github.com/coocood/freecache"
	"github.com/spf13/cast"
	"github.com/spf13/viper"
	"runtime/debug"
)

// @todo Check is freecache without var FreeCache can be used global.
var FreeCache *freecache.Cache

func NewFreeCache() {
	// 50M
	size := cast.ToInt(viper.Get("cache.freecache.size"))
	FreeCache = freecache.NewCache(size)
	// 20%
	debug.SetGCPercent(20)
	return
}
