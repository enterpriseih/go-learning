package mid

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseBetween(head *ListNode, left int, right int) *ListNode {
	// 每次把新加入的节点放在最前面，作为one
	dummyNode := &ListNode{
		Val:  -501,
		Next: head,
	}
	for i := 1; i < left; i++ {
		dummyNode = dummyNode.Next
	}

	cur := dummyNode.Next
	for i := left; i < right; i++ {
		next := cur.Next
		cur.Next = next.Next
		next.Next = dummyNode.Next
		dummyNode.Next = next

		for j := head; j != nil; j = j.Next {
			print(j.Val)
		}
	}

	// 防止第一个节点就需要反转，那么返回的表头应该是dummy的next
	if dummyNode.Val == -501 {
		return dummyNode.Next
	}
	return head
}
