package array

import (
    "fmt"
    "testing"
)

func TestSortArrayByParity(t *testing.T) {
    nums := []int{1,2,3,4}
    re := SortArrayByParity(nums)
    fmt.Println(re)
}

func TestTmp(t *testing.T) {
    nums := []int{10,9,8}
    Tmp(&nums)
    fmt.Println(nums)
}
