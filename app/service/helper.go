package service

import (
	"github.com/gopher-lego/ginger/config"
	"github.com/spf13/cast"
	"math"
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
	URL, err := url.Parse(uri); if err != nil {
		// Parse failure is Invalid
		return ""
	}

	if URL.RawQuery != "" {
		uri = strings.ReplaceAll(uri, URL.RawQuery, "")
	}

	uri = strings.TrimRight(uri, "?")

	return strings.TrimRight(uri, "/")
}

func RightPageNumber(pageNumber int, pageNumberMax int) int {
	var page float64

	page = math.Max(1, cast.ToFloat64(pageNumber))
	page = math.Min(cast.ToFloat64(pageNumberMax), page)

	return cast.ToInt(page)
}