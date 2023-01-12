package silding_window

import (
    "fmt"
    "testing"
)

func TestLengthOfLongestSubstring(t *testing.T) {
    s := "abcbc"
    max := LengthOfLongestSubstring(s)
    fmt.Println(max)
    // fmt.Println(LongestNiceSubstring(s))
}
