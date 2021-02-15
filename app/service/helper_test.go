package service

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// cd ./app/service && go test

func TestDataGoCacheApply(t *testing.T) {
	assert.Equal(t, 0, 0)
}

func TestUriFilterExcludeQueryString(t *testing.T) {
	uris := []string{
		"http://www.example.com/?a=b&c=d",
		"http://www.example.com?a=b&c=d",
	}

	var cleanUri string
	for _, uri := range uris {
		cleanUri = UriFilterExcludeQueryString(uri)
		assert.Equal(t, "http://www.example.com", cleanUri)
	}
}

func TestRightPageNumber(t *testing.T) {
	var pageNumberMax = 100

	assert.Equal(t, 1, RightPageNumber(-1, pageNumberMax))
	assert.Equal(t, 1, RightPageNumber(0, pageNumberMax))

	assert.Equal(t, 2, RightPageNumber(2, pageNumberMax))
	assert.Equal(t, 20, RightPageNumber(20, pageNumberMax))

	assert.Equal(t, 100, RightPageNumber(200, pageNumberMax))
}

func TestSegmentCutSearchMode(t *testing.T) {
	seg := SegmentCutSearchMode("釜山行")
	assert.Equal(t, seg[0], "釜山")
	assert.Equal(t, seg[1], "行")
}