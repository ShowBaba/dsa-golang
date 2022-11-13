package twosum

/*
Given an array of integers nums and an integer target,
return indices of the two numbers such that they add up to target.
*/
func TwoSum(nums []int, target int) [2]int {
	var result = [2]int{}
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if (nums[i] + nums[j]) == target {
				result[0] = i
				result[1] = j
				return result
			}
		}
	}
	return result
}
