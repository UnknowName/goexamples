package main

import "fmt"

func trace(msg string) func() {
	fmt.Println("In trace func ", msg)
	return func() {
		fmt.Printf("IN FUNC'S FUNC")
	}
}

func Debug() {
	// 实际上最后执行的是trace函数里面的函数.第一个print会最开始执行
	// trace func里面的函数打印语句会最后执行，实际上defer是这个函数
	// defer函数中可以获取函数最后返回的值，变可以修改它
	defer trace("debug")()
	fmt.Println("In debug func")
}
