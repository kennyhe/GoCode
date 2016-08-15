package leetcode

func TwoSum(nums [] int, target int)(first, second int) {
	posMap := make(map[int]int)
	for i:=0; i<len(nums); i++ {
		p, ok := posMap[nums[i]]
		if ok {
			first = p
			second = i
			return first, second
		} else {
			posMap[target - nums[i]] = i
		}
	}
	return -1, -1
}