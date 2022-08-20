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
var c map[*common.TreeNode]int
var u map[*common.TreeNode]int

func rob337(root *common.TreeNode) int {
	// 节点被选中c(x)，则左孩子和右孩子都不能被选中
	// 节点没被选中u(x)，则左孩子和右孩子可以被选中，也可以不被选中
	// 所以用c(x) = u(x.l)+u(x.r)
	// u(x) = max(c(x.l), u(x.l)) + max(c(x.r), u(x.r))
	// 后续遍历就是左、右、root
	c = make(map[*common.TreeNode]int)
	u = make(map[*common.TreeNode]int)
	dfsForRob(root)
	return common.GetMax(c[root], u[root])
}

func dfsForRob(root *common.TreeNode) {
	if root == nil {
		return
	}
	dfsForRob(root.Left)
	dfsForRob(root.Right)
	ul := 0
	ur := 0
	cl := 0
	cr := 0
	if root.Left != nil {
		cl = c[root.Left]
		ul = u[root.Left]
	}
	if root.Right != nil {
		cr = c[root.Right]
		ur = u[root.Right]
	}
	c[root] = root.Val + ul + ur
	u[root] = common.GetMax(cl, ul) + common.GetMax(cr, ur)
}
