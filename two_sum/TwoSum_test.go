package twosum

import "testing"

func TestTwoSum(t *testing.T){
	response := TwoSum([]int{2,7,11,15}, 9)
	if response != [2]int{0,1} {
		t.Errorf("error")
	}

	response_ := TwoSum([]int{1,3,5,7,9}, 19)
	if response_ != [2]int{} {
		t.Errorf("error")
	} 
}