package linear_search

import (
	"testing"
)

func TestLinearSearchList(t *testing.T) {
	result := LinearSearch([]int{1, 3, 4, 69, 71, 81, 90, 99, 420, 1337, 69420}, 1336)
	if result {
		t.Errorf("error")
	}
	result = LinearSearch([]int{1, 3, 4, 69, 71, 81, 90, 99, 420, 1337, 69420}, 81)
	if !result {
		t.Errorf("error")
	}

	result = LinearSearch([]int{1, 3, 4, 69, 71, 81, 90, 99, 420, 1337, 69420}, 1)
	if !result {
		t.Errorf("error")
	}
}