package hard

import "fmt"

var queensRes [][]string

func solveNQueens1(n int) [][]string {
	chess := [][]int{}
	for i := 0; i < n; i++ {
		line := []int{}
		for j := 0; j < n; j++ {
			line = append(line, 0)
		}
		chess = append(chess, line)
	}
	queensRes = [][]string{}
	findPos(0, chess)
	return queensRes
}

func findPos(lineNum int, chess [][]int) {
	if lineNum == len(chess) {
		queensRes = append(queensRes, generateChessBoard(chess))
		return
	}
	for j := 0; j < len(chess[lineNum]); j++ {
		if canPlace1(lineNum, j, chess) {
			chess[lineNum][j] = 1
			findPos(lineNum+1, chess)
			chess[lineNum][j] = 0
		}
	}
}

func canPlace1(x, y int, chess [][]int) bool {
	for i := 0; i < len(chess); i++ {
		for j := 0; j < len(chess[i]); j++ {
			if chess[i][j] == 1 {
				if i == x || j == y {
					return false
				}
				if getAbs(x-i) == getAbs(y-j) {
					return false
				}
			}
		}
	}
	return true
}

func getAbs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func generateChessBoard(chess [][]int) (ans []string) {
	for i := 0; i < len(chess); i++ {
		str := ""
		for j := 0; j < len(chess[i]); j++ {
			if chess[i][j] == 1 {
				str += "Q"
			} else {
				str += "."
			}
		}
		ans = append(ans, str)
	}
	return ans
}

func TestSolveNQueens() {
	res := solveNQueens1(1)
	println(fmt.Sprintf("%+v", res))
}
