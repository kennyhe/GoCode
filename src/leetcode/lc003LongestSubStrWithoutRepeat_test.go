package leetcode

import "testing"

func TestLongestSubStrWORepeat1(t *testing.T) {
	str := "abcabcbb"
	lsub := "abc"
	ret := lengthOfLongestSubstring(str)
	if (lsub != ret) {
		t.Errorf("Expected %s, but got %s", lsub, ret)
	}
}


func TestLongestSubStrWORepeat2(t *testing.T) {
	str := "bbbbb"
	lsub := "b"
	ret := lengthOfLongestSubstring(str)
	if (lsub != ret) {
		t.Errorf("Expected %s, but got %s", lsub, ret)
	}
}


func TestLongestSubStrWORepeat3(t *testing.T) {
	str := "pwwkew"
	lsub := "wke"
	ret := lengthOfLongestSubstring(str)
	if (lsub != ret) {
		t.Errorf("Expected %s, but got %s", lsub, ret)
	}
}

