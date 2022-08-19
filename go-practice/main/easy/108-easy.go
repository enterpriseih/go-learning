package easy

import (
	"go-practice/main/common"
)

func sortedArrayToBST(nums []int) *common.TreeNode {
	len := len(nums)
	if len == 0 {
		return nil
	}
	mid := len / 2
	left := sortedArrayToBST(nums[0:mid])
	right := sortedArrayToBST(nums[mid+1 : len])
	root := &common.TreeNode{
		nums[mid],
		left,
		right,
	}
	return root
}
