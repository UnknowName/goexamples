package datastruct

import (
	"errors"
	"fmt"
)

func newCircleQueue(maxsize int) *CircleQueue {
	// 因为使用环形，会导致一个数组的数据无法使用。
	size := maxsize + 1
	return &CircleQueue{
		maxsize: size,
		real:    0,
		front:   0,
		data:    make([]interface{}, size),
	}
}

type CircleQueue struct {
	maxsize int
	// real待进入的队列的元素位置
	real    int
	// front待取出队列的元素位置
	front   int
	data    []interface{}
}

func (q *CircleQueue) String() string {
	// return "CircleQueue{maxsize: " + string(q.maxsize) + "}"
	return fmt.Sprintf(
		"CircleQueue{maxsize: %d, real: %d, front: %d, data: %v}", q.maxsize, q.real, q.front, q.data,
		)
}

func (q *CircleQueue) Put(ele interface{}) error {
	// 先判断当前队列是否为满，为满返回异常
	eleIndex := (q.real + 1) % q.maxsize
	if eleIndex == q.front {
		return errors.New("queue full")
	}
	// 未满时将元素放入，位置标志+1
	q.data[eleIndex] = ele
	q.real ++
	return  nil
}

func (q *CircleQueue) Get() (error, interface{}) {
	// 判断队列是否为空
	if q.real == q.front {
		return errors.New("empty queue"), nil
	}
	eleIndex := (q.front + 1) % q.maxsize
	q.front ++
	return nil, q.data[eleIndex]
}