package repository

import (
	"fmt"
	"os"
	"testing"

	"github.com/gopher-lego/ginger/config"
)

func init()  {
	pwd, _ := os.Getwd()
	config.InitOtherConf(pwd + "/../../setting")
}

// cd ./app/repository && go test
func TestDemo(t *testing.T) {
	fmt.Println(Demo())
}
