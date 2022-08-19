package mid

import (
	"go-practice/main/common"
)

//662
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func widthOfBinaryTree(root *common.TreeNode) int {
	res := 0
	if root == nil {
		return 0
	}
	layer := []*common.TreeNode{}
	index := []int{}

	layer = append(layer, root)
	index = append(index, 0)
	res = 1

	// 找到每一层最左的节点和最右的节点，然后用算法算距离
	for {
		if len(layer) == 0 {
			break
		}
		nextLayer := []*common.TreeNode{}
		nextIndex := []int{}
		for i := 0; i < len(layer); i++ {
			if layer[i].Left != nil {
				nextLayer = append(nextLayer, layer[i].Left)
				nextIndex = append(nextIndex, index[i]*2)
			}
			if layer[i].Right != nil {
				nextLayer = append(nextLayer, layer[i].Right)
				nextIndex = append(nextIndex, index[i]*2+1)
			}
		}
		oneres := 0
		if len(nextIndex) == 1 {
			oneres = 1
		}
		if len(nextIndex) > 1 {
			oneres = nextIndex[len(nextIndex)-1] - nextIndex[0] + 1
		}
		if oneres > res {
			res = oneres
		}
		layer = nextLayer
		index = nextIndex
	}
	return res
}
