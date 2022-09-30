package mid

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func detectCycle(head *ListNode) *ListNode {
	var visited map[*ListNode]int
	node := head
	for {
		if node == nil{
			break
		}
		_, ok := visited[node]
		if ok {
			return node
		} else {
			visited[node] = 1
		}
	}
	return nil
}
