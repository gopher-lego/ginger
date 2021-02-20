package repository

import (
	"fmt"
	"github.com/gopher-lego/ginger/config"
	"os"
	"testing"
)

func init()  {
	pwd, _ := os.Getwd()
	config.InitConf(pwd + "/../../setting")
}

// cd ./app/repository && go test
func TestDemo(t *testing.T) {
	fmt.Println(Demo())
}
