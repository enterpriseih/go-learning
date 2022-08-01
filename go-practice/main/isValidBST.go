package main

//输入：root = [2,1,3]
//输出：true
// -231 <= Node.val <= 231 - 1
func isValidBST(root *TreeNode) bool {
	isValid = true
	hasFirstVal = false
	inorderTraversal(root)
	return isValid
}

var hasFirstVal bool
var pre int
var isValid bool

func inorderTraversal(root *TreeNode) {
	if root == nil || !isValid {
		return
	}
	if root.Left != nil {
		inorderTraversal(root.Left)
	}
	if !hasFirstVal || root.Val > pre {
		pre = root.Val
	} else {
		isValid = false
		return
	}
	if root.Right != nil {
		inorderTraversal(root.Right)
	}
}

// 思路1，一般是左节点<=爸爸<=右节点，可以用递归验证(但要注意左子树如果从右分支过来，要比祖父大；如果从左分支过来要不祖父小)
func isValidBSTWrong(root *TreeNode) bool {
	if root == nil {
		return true
	}
	if root.Left == nil && root.Right == nil {
		return true
	}
	if root.Left != nil {
		if root.Left.Val >= root.Val || !isValidBST(root.Left) {
			return false
		}
	}
	if root.Right != nil {
		if root.Right.Val <= root.Val || !isValidBST(root.Right) {
			return false
		}
	}
	return true
}
