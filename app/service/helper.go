package service

import (
	"github.com/go-ego/gse"
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

// https://github.com/go-ego/gse/blob/master/README_zh.md
func SegmentCutSearchMode(text string) []string {
	var (
		seg gse.Segmenter
	)

	// Skip log print
	seg.SkipLog = true

	// load default dict
	err := seg.LoadDict(); if err != nil {
		panic("segment error")
	}

	// 处理分词结果
	// 支持普通模式和搜索模式两种分词，见代码中 ToString 函数的注释。
	// 搜索模式主要用于给搜索引擎提供尽可能多的关键字
	// seg.String, seg.Slice 输出的类型不同
	return seg.Slice(text, true)
}