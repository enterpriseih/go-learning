package hard

/**
 * Definition for singly-linked list.
*/
type ListNode struct {
     Val int
     Next *ListNode
 }
func reverseKGroup(head *ListNode, k int) *ListNode {
	start := head
	end := head
	for i:=0; i <k-1;i++{
		if end== nil {
			return head
		}
		end = end.Next
	}
	if end== nil {
		return head
	}
	// reverse [start, end]
	// 1\2\3\4
	next :=  start.Next
	start.Next = nil
	for i:=0; i <k-1;i++{
		temp := next.Next
		next.Next = start
		start = next
		next = temp
	}
	head.Next = reverseKGroup(next, k)
	return start
}

func TestReverseKGroup() {
	node5 := &ListNode{
		5,
		nil,
	}
	node4 := &ListNode{
		4,
		node5,
	}
	node3 := &ListNode{
		3,
		node4,
	}
	node2 := &ListNode{
		2,
		node3,
	}
	node1 := &ListNode{
		1,
		node2,
	}

	head := reverseKGroup(node1, 3)
	printList(head)

}

func printList(head *ListNode) {
	temp := head
	for {
		if temp == nil {
			return
		}
		print(temp.Val, " ")
		temp = temp.Next
	}
	println()
}

