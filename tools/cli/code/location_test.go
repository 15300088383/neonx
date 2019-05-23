package code

import (
	"fmt"
	"testing"
)

func TestNewFileLocation(t *testing.T) {

	fmt.Println(
		NewFileLocation("/aaa/bbb/impls/shop/v1/user/login.go", StopImpls))
}

func TestFlushRouterFile(t *testing.T) {

	fmt.Println(LoadModule("/Users/chenhongchu/workspace/code/micro/example/impls/shop/v1/user", StopImpls))
}
