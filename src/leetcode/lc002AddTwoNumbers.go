package leetcode

func AddTwoNumbers(num1, num2 *ListNode) *ListNode {
	if num1 == nil {
		return num2
	}
	if num2 == nil {
		return num1
	}

	fakeHead := new (ListNode)
	rear := fakeHead
	carry := 0
	for ; num1 != nil && num2 != nil; {
		node := new (ListNode)
		sum := num1.data + num2.data + carry
		if (sum >= 10) {
			carry = 1
			sum -= 10
		} else {
			carry = 0
		}
		num1 = num1.next
		num2 = num2.next
		node.data = sum
		rear.next = node
		rear = node
	}

	if (num1 == nil) {
		num1 = num2
	}

	for ; num1 != nil && carry == 1; {
		node := new (ListNode)
		node.data = num1.data + 1
		rear.next = node
		if (node.data == 10) {
			node.data = 0
			carry = 1
			rear = node
			num1 = num1.next
		} else {
			carry = 0
			node.next = num1.next
			break
		}
	}
	if (carry == 1) {
		node := new (ListNode)
		node.data = 1
		rear.next = node
	}
	return fakeHead.next
}