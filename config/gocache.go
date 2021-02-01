package config

import (
	"github.com/patrickmn/go-cache"
	"time"
)

var GoCache *cache.Cache

func NewGoCache() {
	GoCache = cache.New(1 * time.Minute, 30 * time.Minute)
	return
}