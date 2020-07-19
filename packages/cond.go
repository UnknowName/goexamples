package packages

import (
	"fmt"
	"sync"
	"time"
)

func Cond() {
	c := sync.NewCond(&sync.Mutex{})
	queue := make([]interface{}, 0, 10)

	removeFromQueue := func(delay time.Duration) {
		time.Sleep(delay)
		c.L.Lock()
		queue = queue[1:]
		fmt.Println("Removed from queue")
		defer c.L.Unlock()
		c.Signal()
	}

	for i := 0; i < 10; i++ {
		// 当调用wait时，将当前线程阻塞等待
		// 一般用于达到某个条件时主动等待。这里是当队列的长度为2时，主动阻塞
		// c.L.Lock()一定是在条件达到前
		c.L.Lock()
		for len(queue) == 2 {
			c.Wait()
		}
		// 以下代码会因为Wait阻塞，除非条件不成立
		fmt.Println("Adding to queue")
		queue = append(queue, struct{}{})
		go removeFromQueue(1 * time.Second)
		c.L.Unlock()
	}
}