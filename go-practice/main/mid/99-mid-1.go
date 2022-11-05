package mid

import "go-practice/main/common"

func recoverTree1(root *common.TreeNode) {
	// 不排序直接swap
	values1 = []*common.TreeNode{}
	lmr1(root)
	var wrongNode1 *common.TreeNode = nil
	var wrongNode2 *common.TreeNode = nil
	for i := 0; i < len(values1)-1; i++ {
		if values1[i].Val > values1[i+1].Val {
			if wrongNode1 == nil {
				wrongNode1 = values1[i]
			} else {
				wrongNode2 = values1[i+1]
				break
			}
		}
	}
	if wrongNode1 != nil && wrongNode2 != nil {
		temp := wrongNode1.Val
		wrongNode1.Val = wrongNode2.Val
		wrongNode2.Val = temp
	}

}

var values1 []*common.TreeNode

func lmr1(root *common.TreeNode) {
	if root == nil {
		return
	}
	lmr1(root.Left)
	values1 = append(values1, root)
	lmr1(root.Right)
}
