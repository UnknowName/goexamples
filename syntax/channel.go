package main

import (
	"bytes"
	"fmt"
	"net/http"
	"os"
	"sync"
	"time"
)

type Site struct {
	Url    string
	Host   string
	Status int
}

func Check(s *Site) {
	client := &http.Client{Timeout: 5 * time.Second}
	req, _ := http.NewRequest("GET", s.Url, nil)
	req.Host = s.Host
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()
	s.Status = resp.StatusCode
	ch := make(chan *Site)
	// 在往Channel放入数据前，要先准备好从通道取数据的协程，不然主协程会一直阻塞
	// 或者创建channel带缓冲 ch := make(chan *Site, 1)
	go Count(ch)
	ch <- s
	fmt.Println("Check func Over")
}

// 函数声明使用只读Channel。当传入的是双向的时候，函数会进行隐式转换
func Count(ch <-chan *Site) {
	s := <-ch
	fmt.Println("Get site struct from ch", s.Status)
}

// 并发的向同一个Channel发送数据
func ManySend() string {
	responses := make(chan string, 3)
	s := Site{Host: "dev.siss.io", Url: "http://128.0.255.10"}
	// 并发三个请求，放入channel中
	go func() {
		resp, err := http.Get(s.Url)
		if err != nil {
			return
		}
		defer resp.Body.Close()
		responses <- resp.Status
	}()
	go func() {
		resp, err := http.Get(s.Url)
		if err != nil {
			return
		}
		defer resp.Body.Close()
		responses <- resp.Status
	}()
	go func() {
		resp, err := http.Get(s.Url)
		if err != nil {
			return
		}
		defer resp.Body.Close()
		responses <- resp.Status
	}()
	// 从Channel中读取一个数据，这样那个最快哪个将返回
	return <-responses
}

func Wait() {
	fmt.Println("Wait main start")
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		fmt.Println("Gorotinue1 print ")
	}()
	go func() {
		defer wg.Done()
		fmt.Println("Gorotinue2 print")
	}()
	wg.Wait()
	fmt.Println("Wait main over")
}

func ChannelSelect() {
	ch := make(chan int, 1)
	for i := 0; i <= 10; i++ {
		fmt.Println("current i is ", i)
		// select 并不是顺序测试，它会一直等待case语句返回
		select {
		case ch <- i:
		case x := <-ch:
			// 打印出来的是0,2,4,6,8
			// i = 0时，ch <- i执行成功, i = 1时 x := <- ch执行成功。select会尽量保证每个case语句都得到执行
			fmt.Println("x is ", x)
		}
	}
}

func ChannelClose() {
	var stdoutBuf bytes.Buffer
	defer stdoutBuf.WriteTo(os.Stdout)

	intChan := make(chan int, 4)
	go func() {
		defer close(intChan)
		defer fmt.Fprintf(&stdoutBuf, "Done\n")
		for i := 0; i < 5; i++ {
			fmt.Fprintf(&stdoutBuf, "Sending data %d\n", i)
			intChan <- i
		}
		// close(intChan)
	}()

	for i := range intChan {
		fmt.Fprintf(&stdoutBuf, "Received i %d\n", i)
	}
}

func C() {

}
