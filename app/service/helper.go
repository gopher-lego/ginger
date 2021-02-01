package service

import (
	"github.com/gopher-lego/ginger/config"
	"time"
)

func dataCacheApply(key string, data interface{}) {
	expireDuration := time.Duration(60) * time.Minute

	config.GoCache.Set(key, data, expireDuration)

	return
}
