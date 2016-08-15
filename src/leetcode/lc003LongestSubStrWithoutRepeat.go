package leetcode

func lengthOfLongestSubstring(str string) string {
	if len(str) == 0 {
		return str
	}

	lastPos := make([]int, 256)
	for i:=0; i<256; i++ {
		lastPos[i] = -1
	}

	rear := -1
	sublen := 0
	max := 0
	maxSub := ""

	for i:=0; i<len(str); i++ {
		c := str[i]
		if lastPos[c] > rear {
			rear = lastPos[c]
			sublen = i - rear
		} else {
			sublen++
		}
		lastPos[c] = i
		if (sublen > max) {
			max = sublen
			maxSub = str[rear+1:i+1]
		}
	}

	return maxSub
}