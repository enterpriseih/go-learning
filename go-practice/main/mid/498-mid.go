package mid

import "go-practice/main/common"

func findDiagonalOrder(mat [][]int) []int {
	//只有两个方向：右上，x-1，y+1；左下，x+1，y-1
	// x [0, m-1] y [0, n-1]
	// 方向上的数量就是1、2、直到两个值里面的min
	// 总共循环的次数m+n-1
	m := len(mat)
	n := len(mat[0])
	x := 0
	y := 0
	maxLen := common.GetMin(m, n)
	res := []int{}
	res = append(res, mat[x][y])
	for {
		for i := 0; i < maxLen; i++ {
			if isValidPos(x-1, y+1, m, n) {
				x = x - 1
				y = y + 1
				res = append(res, mat[x][y])
			} else {
				break
			}
		}
		if y < n-1 {
			y = y + 1
			res = append(res, mat[x][y])
		} else {
			if x < m-1 {
				x = x + 1
				res = append(res, mat[x][y])
			}
		}
		for i := 0; i < maxLen; i++ {
			if isValidPos(x+1, y-1, m, n) {
				x = x + 1
				y = y - 1
				res = append(res, mat[x][y])
			} else {
				break
			}
		}
		if x < m-1 {
			x = x + 1
			res = append(res, mat[x][y])
		} else {
			if y < n-1 {
				y = y + 1
				res = append(res, mat[x][y])
			}
		}
		if len(res) == m*n {
			break
		}
	}
	return res
}

func isValidPos(x, y, m, n int) bool {
	if x < 0 || x >= m {
		return false
	}
	if y < 0 || y >= n {
		return false
	}
	return true
}
