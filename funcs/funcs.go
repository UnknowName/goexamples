package funcs

import (
	"bytes"
	"fmt"
	"strconv"
)

// Golang的函数不支持关键字参数与默认参数，所有参数必须按定义时的位置传入
func Echo(x, y string) (result string) {
	var buf bytes.Buffer
	buf.WriteString(x)
	buf.WriteString(y)
	result = buf.String()
	return result
}

func Recursive(n int) (total int) {
	/*
	递归函数技巧：
	设定一个退出条件，当达到退出条件时，返回。
	目前实现还有误
	 */
	if n == 1 || n == 0 {
		total = n
	}else{
		total += Recursive(n-1)
	}
	return total
}

func f(i interface{}) (returnStr string) {
	// 将interface类型转换成string类型
	switch i.(type) {
	case string:
		returnStr = i.(string) + " Add string"
	case int:
		returnStr = strconv.Itoa(i.(int)) + " Add string"
	default:
		returnStr = " Add string"
	}
	return returnStr
}

// 函数的参数是函数，可变长参数列表
func FunctionFunction(funcArg func(interface{}) string, args ...interface{}) {
	fmt.Printf("%#v,\n", args)
	fmt.Println(len(args))
	result := funcArg(args[0])
	fmt.Println(result)
}

// 闭包函数
func CloseFunction() (f func() int) {
	var x int
	f = func () int {
		x += 1
		return  x * x
	}
	return f
}