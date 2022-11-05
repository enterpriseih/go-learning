package mid

import "go-practice/main/common"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func kthSmallest(root *common.TreeNode, k int) int {
	// 方案1： 中续遍历，找到第k个
	index = 0
	findK = k
	LMR(root)
	return value

}

var index int
var value int
var findK int

func LMR(root *common.TreeNode) {
	if root == nil {
		return
	}
	LMR(root.Left)
	index++
	if index == findK {
		value = root.Val
	}
	LMR(root.Right)
}
