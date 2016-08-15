package leetcode

import "testing"

func TestIntToList(t *testing.T) {
	num1 := 123
	nums1 := IntToList(num1)
	if nums1.data != 3 {
		t.Errorf("Error in IntToList. Expected 3, but got %d", nums1.data)
	}
	nums1 = nums1.next
	if nums1.data != 2 {
		t.Errorf("Error in IntToList. Expected 2, but got %d", nums1.data)
	}
	nums1 = nums1.next
	if nums1.data != 1 {
		t.Errorf("Error in IntToList. Expected 1, but got %d", nums1.data)
	}
}


func TestListToInt(t *testing.T) {
	num1 := 123
	nums1 := IntToList(num1)
	num1Bak := ListToInt(nums1)
	if num1Bak != num1 {
		t.Errorf("Error in ListToInt. Expected %d, but got %d", num1, num1Bak)
	}
}
