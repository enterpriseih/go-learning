package main

func getMin(x1, x2 int) int {
	if x1 < x2 {
		return x1
	} else {
		return x2
	}
}

//Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
