package mid

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func generateTrees(n int) []*TreeNode {
	return constructTree(1, n)
}

func constructTree(start, end int) []*TreeNode {
	if start == end {
		return []*TreeNode{
			&TreeNode{
				start,
				nil,
				nil,
			},
		}
	}
	if end < start {
		return []*TreeNode{nil}
	}

	res := []*TreeNode{}
	for i := start; i <= end; i++ {
		leftCandidates := constructTree(start, i-1)
		rightCandidates := constructTree(i+1, end)
		for m := 0; m < len(leftCandidates); m++ {
			for n := 0; n < len(rightCandidates); n++ {
				temp := &TreeNode{
					i,
					leftCandidates[m],
					rightCandidates[n],
				}
				res = append(res, temp)
			}
		}
	}
	return res
}
