package array

import (
    "fmt"
    "log"
    "testing"
)

func TestSortArrayByParity(t *testing.T) {
    nums := []int{1,2,3,4}
    re := SortArrayByParity(nums)
    fmt.Println(re)
}

func TestTmp(t *testing.T) {
    nums := []int{10,9,8}
    fmt.Println(nums)
}

func TestFilter(t *testing.T) {
    total := Filter(50)
    log.Println(total)
}
