package mid

func uniquePaths(m int, n int) int {
	mark := [][]int{}
	for i := 0; i < m; i++ {
		mark = append(mark, make([]int, n))
	}
	mark[0][0] = 1
	for j := 0; j < n-1; j++ {
		mark[0][j+1] = mark[0][j]
	}
	for i := 0; i < m-1; i++ {
		mark[i+1][0] = mark[i][0]
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			mark[i][j] = mark[i-1][j] + mark[i][j-1]
		}
	}
	return mark[m-1][n-1]
}

// method 1
//var ans int
//func uniquePaths(m int, n int) int {
//	ans = 0
//	findPath(0, 0, m, n)
//	return ans
//}
//
//func findPath(x, y, m, n int) {
//	if x < 0 || y < 0 || x > m || y > n {
//		return
//	}
//	if x == m-1 && y == n-1 {
//		ans++
//	}
//	findPath(x, y+1, m, n)
//	findPath(x+1, y, m, n)
//
//}

func TestUniquePaths() {
	println(uniquePaths(3, 7))
	println(uniquePaths(3, 2))
	println(uniquePaths(7, 3))
	println(uniquePaths(3, 3))
}
