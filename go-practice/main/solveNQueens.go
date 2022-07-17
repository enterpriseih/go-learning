package main

//51
var num int
var resQueens [][]string

// 一行一行放皇后，每次要找一个可以放皇后的位置（跟已经放置的皇后不同列、不在同一斜线），如果可以找到最后一行【就是一种case】
func solveNQueens(n int) [][]string {
	num = n
	resQueens = [][]string{}
	dfsFindQuene([]int{}, 0)
	return resQueens
}

func dfsFindQuene(queensCols []int, row int) {
	if row == num {
		resQueens = append(resQueens, generateRes(queensCols))
		return
	}
	for j := 0; j < num; j++ {
		if canPlace(queensCols, row, j) {
			queensCols = append(queensCols, j)
			dfsFindQuene(queensCols, row+1)
			//fmt.Printf("%+v, %d, %d\n", queensCols, len(queensCols), row)
			if row > 0 {
				queensCols = queensCols[0:row]
			} else {
				queensCols = []int{}
			}
		}
	}
}

func canPlace(queensCols []int, row int, col int) bool {
	for k := 0; k < row; k++ {
		if col == queensCols[k] {
			return false
		}
		if abs(k, row) == abs(queensCols[k], col) {
			return false
		}
	}
	return true
}

func abs(x1, x2 int) int {
	if x2 > x1 {
		return x2 - x1
	} else {
		return x1 - x2
	}
}

func generateRes(queensCols []int) []string {
	res := []string{}
	for i := 0; i < num; i++ {
		res = append(res, "")
		for j := 0; j < num; j++ {
			if j == queensCols[i] {
				res[i] += "Q"
			} else {
				res[i] += "."
			}
		}
	}
	return res
}
