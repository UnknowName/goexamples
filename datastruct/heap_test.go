package datastruct

import (
    "fmt"
    "testing"
)

func TestNewHeap(t *testing.T) {
    nums := []int{1,2,3,4,5,6,6,7,90}
    heap := NewHeap(nums)
    fmt.Println(heap)
    // heap.Push(10)
    // heap.Push(15)
    // heap.Push(9)
    // heap.Push(91)
    for heap.Size() >= 2 {
        x := heap.Pop()
        y := heap.Pop()
        if x == y {
            continue
        }
        heap.Push(x - y)
    }
    if heap.Size() == 1 {
        v := heap.Pop()
        fmt.Println(v)
    }
    fmt.Println(heap)
}
