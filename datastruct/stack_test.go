package datastruct

import (
	"fmt"
	"testing"
)

func TestNewStack(t *testing.T) {
	stack := NewStack(5)
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	stack.Push(4)
	stack.Push(5)
	// stack.Push("a")
	for i := 0; i < 5; i++ {
		v := stack.Pop()
		fmt.Println(v)
	}
}
