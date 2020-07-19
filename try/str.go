package try

func IsUnique(astr string) bool {
	dic := make(map[rune]struct{})
	for _, v := range astr {
		if _, ok := dic[v]; ok {
			return false
		}
		dic[v] = struct{}{}
	}
	return true
}
