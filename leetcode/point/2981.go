package point

import (
	"log"
)

func MaximumLength(s string) int {
	maxLen := -1
	strLen := len(s)
	substrFreq := make(map[string]int)

	left, right := 0, 0
	for right < strLen {
		for right < strLen && s[right] == s[left] {
			right++
		}
		log.Println("right: ", right)
		for i := left; i < right; i++ {
			subStr := s[left : i+1]
			substrFreq[subStr] += right - i
			log.Println("left", left, "subStr", subStr)
			if substrFreq[subStr] >= 3 && len(subStr) > maxLen {
				maxLen = len(subStr)
			}
		}
		left = right
	}
	return maxLen
}