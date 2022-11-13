package ternarysearch

import "testing"

func TestTenarySearchList(t *testing.T) {
	result := TernarySearch([]int{1, 3, 4, 69, 71, 81, 90, 99, 420, 1337, 69420}, 1336)
	if result {
		t.Errorf("error")
	}
	result = TernarySearch([]int{1, 3, 4, 69, 71, 81, 90, 99, 420, 1337, 69420}, 81)
	if !result {
		t.Errorf("error")
	}

	result = TernarySearch([]int{1, 3, 4, 69, 71, 81, 90, 99, 420, 1337, 69420}, 1)
	if !result {
		t.Errorf("error")
	}
}