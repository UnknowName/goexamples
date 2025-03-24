package structs

import (
    "container/heap"
    "fmt"
    "testing"
)

func TestIntHeap_Less(t *testing.T) {
    nums := &IntHeap{1,20,3,5,-1}
    heap.Init(nums)
    heap.Push(nums, -2)
    fmt.Println(nums)
    for i := 0; i < 6; i++ {
        v := heap.Pop(nums)
        fmt.Println(v)
    }
}
