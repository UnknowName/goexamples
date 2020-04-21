package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

/*
func main() {
	fmt.Println("Hello, world!")
	//fmt.Printf函数最后不会自动加上换行符
	fmt.Printf("%s, go!\n", "Hello")
}
 */

func main() {
	host := os.Args[1]
	url := "https://shopapi.sissyun.com.cn/home/getQueueIsAbnormal?host=" + host
	fmt.Println(url)
	resp, err := http.Get(url)
	if err != nil {
		// fmt.Println("Debug ", err)
		fmt.Println(1)
		return
	}
	defer resp.Body.Close()
	byteData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// fmt.Println("Debug 2 ", err)
		fmt.Println(1)
		return
	}
	if resp.StatusCode != http.StatusOK {
		fmt.Println(resp.StatusCode)
		return
	}
	fmt.Println(string(byteData))
}
