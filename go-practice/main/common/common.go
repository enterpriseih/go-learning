package common

import "fmt"

func GetMin(x1, x2 int) int {
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

func PrintSlice(x []int) {
	fmt.Printf("len=%d cap=%d slice=%v\n", len(x), cap(x), x)
}
