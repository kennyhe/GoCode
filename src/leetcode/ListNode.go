package leetcode

type ListNode struct {
	data int
	next *ListNode
}

func IntToList(num int) *ListNode {
	if num <= 0 {
		node := new (ListNode)
		node.data = 0
		return node
	}

	fakeHead := new(ListNode)
	rear := fakeHead
	for ; num > 0; {
		node := new (ListNode)
		node.data = num % 10
		num /= 10
		rear.next = node
		rear = node
	}

	return fakeHead.next
}

func ListToInt(list *ListNode) int {
	if list == nil {
		return -1
	}

	result := 0
	w := 1
	for ; list != nil; {
		result += (list.data * w)
		w *= 10
		list = list.next
	}
	return result
}