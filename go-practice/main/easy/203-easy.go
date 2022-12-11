package easy

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeElements(head *ListNode, val int) *ListNode {
	temp := head
	// 找到第一个不是val, 新head
	for {
		if temp == nil {
			return nil
		}
		if temp.Val != val {
			break
		}
		temp = temp.Next
	}
	head = temp
	pre := temp
	for {
		temp = temp.Next
		if temp == nil {
			break
		}
		if temp.Val == val {
			pre.Next = temp.Next
		} else {
			pre = temp
		}
	}
	return head
}
