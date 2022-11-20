package _120

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func closestNodes(root *TreeNode, queries []int) [][]int {
	var res [][]int
	maxTree = -1
	minTree = -1
	root = balanceBST(root)
	traversal(root)
	for i := 0; i < len(queries); i++ {
		a := -1
		b := -1
		var oneres []int
		if queries[i] == minTree || queries[i] == maxTree {
			a = queries[i]
			b = queries[i]
		}
		if queries[i] > minTree {
			a = findLeftClosest(root, queries[i])
		}
		if queries[i] < maxTree {
			b = findRightClosest(root, queries[i])
		}
		oneres = []int{a, b}
		res = append(res, oneres)
	}
	return res
}

func findLeftClosest(root *TreeNode, key int) int {
	if root == nil {
		return -1
	}
	if root.Val == key {
		return key
	}
	if root.Val < key {
		return max(findLeftClosest(root.Right, key), root.Val)
	}
	if root.Val > key {
		return findLeftClosest(root.Left, key)
	}
	return -1
}

func findRightClosest(root *TreeNode, key int) int {
	if root == nil {
		return -1
	}
	if root.Val == key {
		return key
	}
	if root.Val < key {
		return findRightClosest(root.Right, key)
	}
	if root.Val > key {
		res := findRightClosest(root.Left, key)
		if res == -1 {
			return root.Val
		}
		return min(res, root.Val)
	}
	return -1
}

var maxTree int
var minTree int

func traversal(root *TreeNode) {
	if root == nil {
		return
	}
	if maxTree == -1 || root.Val > maxTree {
		maxTree = root.Val
	}
	if minTree == -1 || root.Val < minTree {
		minTree = root.Val
	}
	traversal(root.Left)
	traversal(root.Right)
}
func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}

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

func TestClosestNodes() {
	a := &TreeNode{
		9,
		nil,
		nil,
	}

	b := &TreeNode{
		4,
		nil,
		a,
	}
	println(fmt.Sprintf("%+v", closestNodes(b, []int{3})))
}
