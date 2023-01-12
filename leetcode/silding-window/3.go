package silding_window

func LengthOfLongestSubstring(s string) int {
    // 滑动窗口 + 字典
    if len(s) <= 1 {
        return len(s)
    }
    dict := map[byte]struct{}{}
    left := 0
    right := 0
    max := 0
    // 关键是先走left
    for left < len(s) && right < len(s) {
        // 如果当前字符在之前没出现过，窗口扩大
        if _, ok := dict[s[right]]; !ok {
            dict[s[right]] = struct{}{}
            right += 1
        } else {
            // 说明之前的字符在当前窗口出现过，先把最左的字符出窗口，再将左窗口右移
            delete(dict, s[left])
            left += 1
        }
        // 取最大值
        if max < right - left  {
            max = right - left
        }
    }
    return max
}