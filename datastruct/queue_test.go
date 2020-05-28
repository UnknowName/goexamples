package datastruct

import (
	"fmt"
	"testing"
)

func TestCircleQueue(t *testing.T) {
	queue := newCircleQueue(3)
	fmt.Println(queue.Get())
	for i := 0; i < 3; i++ {
		if err := queue.Put(i); err != nil {
			fmt.Println(err)
			break
		}
	}
	// 已满，不能再放了
	fmt.Println(queue.Put("a"))
	// 取出元素，依次为0，1，2
	for i := 0; i < 3; i ++ {
		fmt.Println(queue.Get())
	}
	// 又可以放了
	fmt.Println(queue.Put("a"))
	//取出来
	fmt.Println(queue.Get())
	// 再取为空
	fmt.Println(queue.Get())
}
