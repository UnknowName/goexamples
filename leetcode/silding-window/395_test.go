package silding_window

import (
	"fmt"
	"testing"
)

func TestLongestSubstring(t *testing.T) {
    s := "aaabb"
    LongestSubstring(s, 2)
}


func TestMax(t *testing.T) {
    s := "aeu"
    v := MaxVowels(s, 2)
    fmt.Println(v)
}

func TestChangeNumber(t *testing.T) {
    v := divisorSubstrings(430043, 2)
    fmt.Println(v)
}