package main

import (
	"fmt"
	"testing"
	"time"
)

func TestCheck(t *testing.T) {
	s := &Site{
		Url:    "http://128.0.255.254/demo/daohao",
		Host:   "dev.siss.io",
		Status: 0,
	}
	Check(s)
	// 主进程要稍等一会儿
	time.Sleep(time.Second)
}

func TestManySend(t *testing.T) {
	status := ManySend()
	fmt.Println(status)
}

func TestWait(t *testing.T) {
	Wait()
}

func TestChannelSelect(t *testing.T) {
	ChannelSelect()
}

func TestChannelClose(t *testing.T) {
	ChannelClose()
}

func TestC(t *testing.T) {
	doWork := func(done <-chan interface{}, strings chan string) <-chan string {
		returnStream := make(chan string)
		go func() {
			defer fmt.Println("doWork func exit")
			defer close(returnStream)
			for {
				select {
				case s := <-strings:
					returnStream <- s
				case <-done:
					return
				}
			}
		}()
		return returnStream
	}

	done := make(chan interface{}, 1)
	stringChan := doWork(done, nil)
	go func(){
		defer close(done)
		time.Sleep(time.Second)
		fmt.Println("Close done chan and cancel goroutine")
	}()

	for s := range stringChan {
		fmt.Println("Received  string is ", s)
	}
}

func TestCount(t *testing.T) {
	longStream := func(done <-chan interface{}) <-chan int {
		intStream := make(chan int)
		go func() {
			defer fmt.Print("Exit go func")
			defer close(intStream)
			for i := 1; i <= 5; i++ {
				time.Sleep(time.Second)
				intStream <- i
			}
		}()
		return intStream
	}
	done := make(chan interface{})
	intChan := longStream(done)
	for num := range intChan {
		fmt.Print(num)
	}
	close(done)
}

func TestC2(t *testing.T) {
	f := func(values []int, leng int) []int {
		result := make([]int, len(values))
		for i, v := range values {
			result[i] = v * i
		}
		return result
	}

	nums := []int{0, 2, 3, 4}
	newNums := f(nums, len(nums))
	fmt.Print(newNums)
}