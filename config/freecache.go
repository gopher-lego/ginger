package config

import (
	"github.com/coocood/freecache"
	"github.com/spf13/cast"
	"runtime/debug"
)

var FreeCache *freecache.Cache

func NewFreeCache() {
	// 50M
	size := cast.ToInt(Get("cache.freecache.size"))
	FreeCache = freecache.NewCache(size)
	// 20%
	debug.SetGCPercent(20)
	return
}
