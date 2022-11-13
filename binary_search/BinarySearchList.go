package binary_search

func BinarySearch(haystack []int, needle int) bool {
	lo := 0
	hi := len(haystack)
	for {
		m := lo + (hi-lo)/2 // get mid-point
		v := haystack[m]
		if v == needle {
			return true
		} else if v > needle {
			hi = m
		} else {
			lo = m + 1
		}
		if !(lo < hi) {
			break
		}
	}
	return false
}
