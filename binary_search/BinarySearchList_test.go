package binary_search

import (
	"testing"
)

func TestBinarySearchList(t *testing.T) {
	result := BinarySearch([]int{1, 3, 4, 69, 71, 81, 90, 99, 420, 1337, 69420}, 1336)
	if result {
		t.Errorf("error")
	}
	result = BinarySearch([]int{1, 3, 4, 69, 71, 81, 90, 99, 420, 1337, 69420}, 81)
	if !result {
		t.Errorf("error")
	}

	result = BinarySearch([]int{1, 3, 4, 69, 71, 81, 90, 99, 420, 1337, 69420}, 1)
	if !result {
		t.Errorf("error")
	}
}