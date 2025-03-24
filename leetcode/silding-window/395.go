package silding_window

func LongestSubstring(s string, k int) {

}

func findAnagrams(s string, p string) []int {
	res := make([]int, 0)
	if len(p) > len(s) {
		return res
	}
	count1 := [26]int{}
	count2 := [26]int{}
	for i := range p {
		index := p[i] - 'a'
		count1[index] += 1
	}
	left := 0
	right := 0
	for right < len(s) {
		// 因为搜索子中，所以窗口大小就是子串的长度
		for right < len(p) {
			index := s[right] - 'a'
			count2[index]++
			right++
		}
		if count1 == count2 {
			res = append(res, left)
		}
		if left < len(s) && right < len(s) {
			count2[s[right]-'a'] += 1
			count2[s[left]-'a'] -= 1
		}
		left += 1
		right += 1
	}
	return res
}

func findAnagramsOk(s string, p string) []int {
	totals := make([]int, 0)
	if len(p) > len(s) {
		return totals
	}
	// 改成用数组，效率提升好多
	count1 := [26]int{}
	count2 := [26]int{}
	for i := range p {
		index := p[i] - 'a'
		count1[index] += 1
	}
	left := 0
	for i := 0; i <= len(s); i++ {
		for i-left < len(p) {
			index := s[i] - 'a'
			count2[index] += 1
			i++
			continue
		}
		if count1 == count2 {
			totals = append(totals, left)
		}
		if left < len(s) && i < len(s) {
			count2[s[i]-'a'] += 1
			count2[s[left]-'a'] -= 1
		}
		left++
	}
	return totals
}

func MaxVowels(s string, k int) int {
	base := map[byte]bool{'a': true, 'e': true, 'i': true, 'o': true, 'u': true}
	left := 0
	re := 0
	max := 0
	for i := 0; i < len(s); i++ {
		if base[s[i]] {
			re++
		}
		if i < k - 1 {
			continue
		}
        if max < re {
            max = re
        }
        if base[s[left]] {
            re--
        }
		left++
	}
	return max
}
