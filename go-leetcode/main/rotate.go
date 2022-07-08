package main

// 思路：先去对角，然后横着翻转
func rotate(matrix [][]int) {
	n := len(matrix)
	// 对角线翻转，就是swap(matrix[i][j], matrix[j][i])
	for i := 1; i < n; i++ {
		for j := 0; j < i; j++ {
			temp := matrix[i][j]
			matrix[i][j] = matrix[j][i]
			matrix[j][i] = temp
		}
	}
	// 横着翻转
	for i := 0; i < n; i++ {
		for j := 0; j < n/2; j++ {
			temp := matrix[i][j]
			matrix[i][j] = matrix[i][n-1-j]
			matrix[i][n-1-j] = temp
		}
	}
}
