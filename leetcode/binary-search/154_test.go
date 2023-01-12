package binary_search

import (
    "log"
    "testing"
)

func TestFindMin(t *testing.T) {
    nums := []int{10,1,10,10,10}
    x := FindMin(nums)
    log.Println("x = ", x)
}
