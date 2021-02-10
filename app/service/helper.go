package service

import (
	"github.com/gopher-lego/ginger/config"
	"net/url"
	"strings"
	"time"
)

func DataGoCacheApply(key string, data interface{}) {
	expireDuration := time.Duration(60) * time.Minute

	config.GoCache.Set(key, data, expireDuration)

	return
}

func UriFilterExcludeQueryString(uri string) string {
	URL, _ := url.Parse(uri)

	if URL.RawQuery != "" {
		uri = strings.ReplaceAll(uri, URL.RawQuery, "")
	}

	uri = strings.TrimRight(uri, "?")

	return strings.TrimRight(uri, "/")
}