package mid

//type TreeNode struct {
//	Val   int
//	Left  *TreeNode
//	Right *TreeNode
//}

func balanceBST(root *TreeNode) *TreeNode {
	var values []int
	var traversalGetV func(*TreeNode)
	traversalGetV = func(node *TreeNode) {
		if node == nil {
			return
		}
		traversalGetV(node.Left)
		values = append(values, node.Val)
		traversalGetV(node.Right)
	}
	traversalGetV(root)

	var buildTree func([]int) *TreeNode
	buildTree = func(values []int) *TreeNode {
		if len(values) == 0 {
			return nil
		}
		mid := len(values) / 2
		root := &TreeNode{
			values[mid],
			nil,
			nil,
		}
		if mid > 0 {
			root.Left = buildTree(values[0:mid])
		}
		if (mid + 1) < len(values) {
			root.Right = buildTree(values[mid+1:])
		}
		return root
	}
	return buildTree(values)
}
