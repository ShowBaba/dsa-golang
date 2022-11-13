package ternarysearch

func TernarySearch(haystack []int, needle int) bool {
	lo := 0
	hi := len(haystack)
	for {
		m1 := lo + (hi-lo)/3
		m2 := hi - (hi-lo)/3

		v1 := haystack[m1]
		v2 := haystack[m2]

		if v1 == needle || v2 == needle {
			return true
		} else if v1 > needle {
			hi = m1
		} else if v2 > needle {
			hi = m2
			lo = m1 + 1
		} else {
			lo = m2
		}
		if !(lo < hi) {
			break
		}
	}
	return false
}
