package competition

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

var mark map[int]int

func numColor(root *TreeNode) int {
	mark = make(map[int]int)
	traverse(root)
	return len(mark)
}

func traverse(root *TreeNode) {
	if root == nil {
		return
	}
	_, ok := mark[root.Val]
	if !ok {
		mark[root.Val]=1
	}
	traverse(root.Left)
	traverse(root.Right)
}