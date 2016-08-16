package leetcode

import "testing"

func TestIsPalindrome(t *testing.T) {
	s := "abcd"
	if isPalindrome(s, 0, len(s)-1) {
		t.Errorf("%s is not a palindrome, but your code make it is", s)
	}
	s1 := "abcdcba"
	if ! isPalindrome(s1, 0, len(s1)-1) {
		t.Errorf("%s is a palindrome, but your code make it is not", s1)
	}
}


func TestLongestPalindromeSubStr1(t *testing.T) {
	str := "abcabcbb"
	lsub := "bcb"
	ret := LongestPalindromeSubStr(str)
	if (lsub != ret) {
		t.Errorf("Expected %s, but got %s", lsub, ret)
	}
}
