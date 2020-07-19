package packages

import (
	"context"
	"fmt"
	"net/http"
	"sync"
	"time"
)

/*
 Context用于跟踪Goroutine链的，比如一个Goroutine又启动了很多Goroutine，而这些都要被取消
 但需要配合Select语法一起使用
*/

// 相关子Goroutine取消后全部取消
func Context() {
	go http.ListenAndServe(":8099", nil)
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		time.Sleep(time.Second * 3)
		cancel()
	}()
	fmt.Println(A(ctx))
	select {}
}

func C(ctx context.Context) string {
	select {
	case <-ctx.Done():
		reason := ctx.Err()
		if reason != nil {
			s := reason.Error()
			return "C Done " + s
		}
		return "C Done"
	}
}

func B(ctx context.Context) string {
	go fmt.Println(C(ctx))
	select {
	case <-ctx.Done():
		return "B Done"
	}
}

func A(ctx context.Context) string {
	go fmt.Println(B(ctx))
	select {
	case <-ctx.Done():
		return "A Done"
	}
}

// Goroutine超时取消
func DeadContext() {
	timeout := time.Now().Add(time.Second * 3)
	// ctx设置为3秒超时.WithTime实际上就是WithDeadline
	ctx, cancel := context.WithDeadline(context.Background(), timeout)
	//ctx, cancel := context.WithTimeout(context.Background(), 3 * time.Second)
	var wg sync.WaitGroup
	wg.Add(1)
	go func(ctx context.Context) {
		defer wg.Done()
		select {
		case <-time.After(time.Second * 2):
			// 这里2秒后就会退出，因为下面的ctx.Done()代码快会跳过
			fmt.Println("After 2 second, the goroutine exit")
			fmt.Println(ctx.Err())
			fmt.Println(ctx.Deadline())
		case <-ctx.Done():
			// 实际这里可以用于等待一个goroutine。
			fmt.Println("Goroutine timeout end")
			fmt.Println(ctx.Err())
		}
	}(ctx)
	// 取消
	cancel()
	wg.Wait()
}

// 可以简单的看起是一个map类型，用于在Goroutine间传递的map
func ValueContext() {
	// 这里故意将一般类型转换为指定类型，是因为Value的参数是Interface。防止重复的数据被弄乱，这样传进去的是不同类型，即使
	// 相同值的变量也会被区分开来
	type userID string
	type token string
	handResponse := func(ctx context.Context) {
		fmt.Println(ctx.Value("userId"))
		fmt.Println(ctx.Value("token"))
	}
	processRequest := func(userId userID, authToken token) {
		ctx := context.WithValue(context.Background(), "userId", userId)
		ctx = context.WithValue(ctx, "token", authToken)
		handResponse(ctx)
	}
	processRequest("1001", "abc")
}