package _007

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func expandBinaryTree(root *TreeNode) *TreeNode {
	traverse(root)
	return root
}

func traverse(root *TreeNode) {
	if root == nil {
		return
	}
	if root.Left != nil {
		insert1 := &TreeNode{
			-1,
			root.Left,
			nil,
		}
		root.Left = insert1
		traverse(insert1.Left)
	}
	if root.Right != nil {
		insert2 := &TreeNode{
			-1,
			nil,
			root.Right,
		}
		root.Right = insert2
		traverse(insert2.Right)
	}
}

func printTree(root *TreeNode) {
	if root == nil {
		return
	}
	print(root.Val)
	printTree(root.Left)
	printTree(root.Right)
}
func TestExpandBinaryTree() {
	left := &TreeNode{
		5,
		nil,
		nil,
	}
	right := &TreeNode{
		6,
		nil,
		nil,
	}
	root := &TreeNode{
		7,
		left,
		right,
	}
	printTree(expandBinaryTree(root))
}
