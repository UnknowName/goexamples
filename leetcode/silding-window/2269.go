package silding_window

import (
    "strconv"
)

func divisorSubstrings(num int, k int) int {
    s := strconv.Itoa(num)
    re := 0
    left := 0
    for i := 0; i < len(s); i++ {
        if i < k - 1   {
            continue
        }
        v := ChangeNumber(s[left:i+1])
        // fmt.Println("v=", v)
        if v > 0 && num % v == 0 {
            re++
        }
        left++
    }
    return re
}

func ChangeNumber(s string) int {
    sum := 0
    base := 1
    for i := len(s) - 1; i >= 0; i-- {
        num := s[i] - '0'
        sum += base * int(num)
        base *= 10
    }
    return sum
}
