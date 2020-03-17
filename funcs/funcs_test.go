package funcs

import (
	"fmt"
	"testing"
)

func TestEcho(t *testing.T) {
	total := Echo("hello ", "word")
	fmt.Println(total)
}

func TestRecursive(t *testing.T) {
	fmt.Println(Recursive(3))
}

func TestFunctionFunction(t *testing.T) {
	// 多个参数传递时，直接传
	FunctionFunction(f, 12, "bbbb")
	// 如果原始参数就是切片类型，参数传时增加省略号
	funcArg := []interface{}{"test", "aaa"}
	FunctionFunction(f, funcArg...)
}

func TestCloseFunction(t *testing.T) {
	f := CloseFunction()
	fmt.Println(f())
	fmt.Println(f())
}