package http

import (
	"os"
	"testing"

	"github.com/gopher-lego/ginger/config"
)

func init()  {
	pwd, _ := os.Getwd()
	config.InitOtherConf(pwd + "/../../setting")
}

func TestPong(t *testing.T) {
}