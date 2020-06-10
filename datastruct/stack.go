package datastruct

import (
	"fmt"
	originSync "sync"
)

// 加入RWMutex保证线程安全
type Stack struct {
	data []interface{}
	max  int
	top  int
	lock originSync.RWMutex
}

func NewStack(max int) *Stack {
	return &Stack{
		data: make([]interface{}, max),
		max:  max,
		top:  -1,
		lock: originSync.RWMutex{},
	}
}

func (s *Stack) Empty() bool {
	return s.top == -1
}

func (s *Stack) Full() bool {
	return s.top == s.max - 1
}

// 出栈
func (s *Stack) Pop() interface{} {
	// 先判断是否为空
	if s.Empty() {
		panic("当前栈空，不能出栈")
	}
	value := s.data[s.top]
	s.lock.Lock()
	s.top--
	defer s.lock.Unlock()
	return value
}

// 入栈
func (s *Stack) Push(value interface{}) {
	if s.Full() {
		panic("栈满，不能入栈")
	}
	s.lock.Lock()
	s.top++
	s.data[s.top] = value
	defer s.lock.Unlock()
}

func (s *Stack) Show() {
	for i := s.top; i >= 0; i-- {
		// dataType := reflect.TypeOf(s.data[i])
		switch s.data[i].(type) {
		case int:
			fmt.Println(s.data[i])
		case int32:
			fmt.Println(string(s.data[i].(int32)))
		}
	}
}

func (s *Stack) Top() interface{} {
	return s.data[s.top]
}