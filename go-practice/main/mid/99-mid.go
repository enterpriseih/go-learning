package mid

import (
	"go-practice/main/common"
	"sort"
)

func recoverTree(root *common.TreeNode) {
	// 基本的还是很简单的, 排序然后中序遍历赋值
	values = []int{}
	i = 0
	lmr(root)
	sort.Slice(values, func(i, j int) bool {
		if values[i] < values[j] {
			return true
		} else {
			return false
		}
	})
	recover(root)

}

var values []int
var i int

func lmr(root *common.TreeNode) {
	if root == nil {
		return
	}
	lmr(root.Left)
	values = append(values, root.Val)
	lmr(root.Right)
}

func recover(root *common.TreeNode) {
	if root == nil {
		return
	}
	recover(root.Left)
	root.Val = values[i]
	i++
	recover(root.Right)
}
