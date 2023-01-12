package dp

import "fmt"

func MinCostClimbingStairs(cost []int) int {
    if len(cost) <= 1 {
        return 0
    }
    dp := make([]int, len(cost))
    dp[0] = cost[0]
    dp[1] = cost[1]
    for i := 2; i < len(cost); i++ {
        dp[i] = min(dp[i - 1] + cost[i], dp[i - 2] + cost[i])
        fmt.Println(dp, cost[i])
    }
    return min(dp[len(dp) - 1], dp[len(dp) - 2])
}

func min(n1, n2 int) int {
    if n1 < n2 {
        return n1
    }
    return n2
}
