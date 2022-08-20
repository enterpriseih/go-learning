package mid

import "fmt"

/**
* Definition for a Node.
 */
type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	// 创建一个map，索引是原链表的地址，值是新链表的值
	if head == nil {
		return nil
	}
	mark := map[*Node]*Node{}
	temp := head
	newHead := &Node{
		Val:    temp.Val,
		Next:   nil,
		Random: nil,
	}
	pre := newHead
	for {
		mark[temp] = pre
		if temp.Next == nil {
			break
		}
		newNode := &Node{
			Val:    temp.Next.Val,
			Next:   nil,
			Random: nil,
		}
		pre.Next = newNode
		pre = newNode
		temp = temp.Next
	}
	temp = head
	for {
		if temp == nil {
			break
		}
		if temp.Random == nil {
			temp = temp.Next
			continue
		}
		curRandom := temp.Random
		newTemp := mark[temp]
		newRandom := mark[curRandom]
		newTemp.Random = newRandom

		temp = temp.Next
	}
	return newHead
}

func TestCopyRandomList() {
	a := &Node{1, nil, nil}
	b := &Node{2, nil, nil}
	c := &Node{3, nil, nil}
	a.Next = b
	b.Next = c
	a.Random = c
	b.Random = a
	c.Random = a
	x := copyRandomList(a)
	for {

		if x.Next == nil {
			break
		}
		fmt.Printf("val=%d, nextVal=%d, randomVal=%d\n", x.Val, x.Next.Val, x.Random.Val)
		x = x.Next
	}

}
