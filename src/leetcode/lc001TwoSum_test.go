package leetcode

import "testing"

func TestTwoSum(t *testing.T) {
	nums := []int {1,2,8, 13, 24}
	target := 14
	pos1, pos2 := TwoSum(nums, target)
	if pos1 != 0 || pos2 != 3 {
		t.Errorf("Expected 1,3, but got %d, %d", pos1, pos2)
	}

	target = 21
	pos1, pos2 = TwoSum(nums, target)
	if pos1 != 2 || pos2 != 3 {
		t.Errorf("Expected 2,3, but got %d, %d", pos1, pos2)
	}
}