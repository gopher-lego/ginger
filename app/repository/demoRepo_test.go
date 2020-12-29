package repository

import (
	"fmt"
	"testing"
)

// cd ./app/repository && go test
func TestDemoRepo(t *testing.T) {
	fmt.Println(DemoRepo())
}
