package leetcode

import "testing"

func TestAddTwoNumbers(t *testing.T) {
	num1 := 123
	nums1 := IntToList(num1)

	num2 := 456
	nums2 := IntToList(num2)

	expectedSum := 579

	sumList := AddTwoNumbers(nums1, nums2)
	sum := ListToInt(sumList)
	if sum != expectedSum {
		t.Errorf("Expected %d, but got %d", expectedSum, sum)
	}
}


func TestAddTwoNumbers2(t *testing.T) {
	num1 := 99
	nums1 := IntToList(num1)

	num2 := 99999
	nums2 := IntToList(num2)

	expectedSum := 100098

	sumList := AddTwoNumbers(nums1, nums2)
	sum := ListToInt(sumList)
	if sum != expectedSum {
		t.Errorf("Expected %d, but got %d", expectedSum, sum)
	}
}
