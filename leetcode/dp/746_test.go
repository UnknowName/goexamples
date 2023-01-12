package dp

import (
    "fmt"
    "testing"
)

func TestMinCostClimbingStairs(t *testing.T) {
    nums := []int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1}
    fmt.Println(MinCostClimbingStairs(nums))
}
