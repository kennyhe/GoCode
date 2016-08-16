package leetcode

func LongestPalindromeSubStr(s string) string {
	if len(s) <= 1 {
		return s
	}

	maxLen := 1
	maxSub := s[0:1]
	for pos:=1; pos<len(s); pos++ {
		if pos - maxLen >= 1 && isPalindrome(s, pos - maxLen - 1, pos) {
			maxLen += 2
			maxSub = s[pos-maxLen+1:pos+1]
		} else if isPalindrome(s, pos-maxLen, pos) {
			maxLen++
			maxSub = s[pos-maxLen+1:pos+1]
		}
	}
	return maxSub
}


func isPalindrome(s string, start, end int) bool {
	for  ; start < end;  {
		if s[start] != s[end] {
			return false
		}
		start++
		end--
	}
	return true
}