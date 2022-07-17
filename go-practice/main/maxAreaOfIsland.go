package main

// 695
// dfs 如果走到水坑就不继续往下；通过dfs找寻一颗陆地连通树，面积就是树节点的sum
var m int
var n int
var maxArea int

func maxAreaOfIsland(grid [][]int) int {
	m = len(grid)
	n = len(grid[0])
	maxArea = 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			area := dfs(grid, i, j)
			if area > maxArea {
				maxArea = area
			}
		}
	}
	return maxArea
}

func dfs(grid [][]int, i int, j int) int {
	if i < 0 || i >= m || j < 0 || j >= n || grid[i][j] == 0 {
		return 0
	}
	// 经过一块土地就将其置为0，这样不会重复计算
	grid[i][j] = 0
	res := 1
	// 上：i-1,j
	// 下：i+1,j
	// 左: i, j-1
	// 右: i, j+1
	res += dfs(grid, i-1, j)
	res += dfs(grid, i+1, j)
	res += dfs(grid, i, j-1)
	res += dfs(grid, i, j+1)
	return res
}
