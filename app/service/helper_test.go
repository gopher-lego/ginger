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