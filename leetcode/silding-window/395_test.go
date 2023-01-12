package silding_window

import (
    "fmt"
    "testing"
)

func TestLongestSubstring(t *testing.T) {
    strs := []string{"bbbdd"}
    for _, str := range strs {
        re := LongestSubstring(str, 2)
        fmt.Println(re)
    }
}
