package http

import (
	"github.com/gopher-lego/skeleton/config"
	"os"
	"testing"
)

func init()  {
	pwd, _ := os.Getwd()
	config.InitConf(pwd + "/../../setting")
}

func TestPong(t *testing.T) {
}