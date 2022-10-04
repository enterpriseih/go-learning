package competition

import (
	"fmt"
	"sort"
)

func circleGame(toys [][]int, circles [][]int, r0 int) (ans int) {
	sort.Slice(circles, func(i, j int) bool { a, b := circles[i], circles[j]; return a[0] < b[0] || a[0] == b[0] && a[1] < b[1] })
	//可以先换成一个简单的问题，一个圈和一个玩具，如何判断玩具在圈内
	// 优化1：注意要先遍历toy，再遍历circle，因为toy如果判断在圈内就可以break
	// 优化2：不需要遍历所有的circle，因为半径最大也是11，可以找toy圆心附近的最大circle 【但是二分查找只能用一个纬度，两个纬度找不到】
	// 优化3： circle排序之后有重复的，可以删除那些重复的
	println(fmt.Sprintf("%v", circles))
	ans = 0
	for j, _ := range toys {

		toyX := toys[j][0]
		toyY := toys[j][1]
		//println(fmt.Sprintf("%v", circles))
		println(toyX, toyY)
		start := sort.Search(len(circles), func(i int) bool { return circles[i][0] > toyX-11})
		end := sort.Search(len(circles), func(i int) bool { return circles[i][0] > toyX+11})
		println(start, end)
		for ; start <=end && start<len(circles); start++ {
			if judgeToyInCircle(toys[j], circles[start], r0) {
				ans++
				println(fmt.Sprintf("%v", toys[j]))
				break
			}
		}
	}


	return ans
}


func judgeToyInCircle(toy []int, circle []int, r int) bool{
	if toy[2] > r {
		return false
	}
	if (int64)((r-toy[2])*(r-toy[2])) >= disSquare(toy[0], toy[1], circle[0], circle[1]) {
		return true
	}
	return false
}


func  disSquare(x1,y1, x2,y2 int) int64 {
	return (int64)((x2-x1)*(x2-x1) +(y2-y1)*(y2-y1))
}

func TestCircleGame() {
	//[[3,3,1],[3,2,1]], circles = [[4,3]], r = 2
//[[1,3,2],[4,3,1],[7,1,2]], circles = [[1,0],[3,3]], r = 4
	println(circleGame([][]int{{1,23,2}},
		[][]int{{75,84},{70,20},{46,41},{7,37},{32,70},{56,38},{67,56},{48,53},{56,21},{8,96},{39,81},{79,59},{64,25},{26,96},{70,38},{33,55},{54,77},
			{70,74},{73,30},{53,41},{88,82},{47,87},{77,20},{55,38},{12,91},{50,97},{90,6},{27,34},{7,74},{80,56},{76,50},{83,79},{2,86},{27,9},{46,95},
			{48,50},{27,21},{15,46},{59,6},{23,34},{50,89},{94,20},{76,66},{50,90},{24,13},{31,5},{48,71},{71,89},{82,7},{80,82},{53,78},{83,86},{82,10},
			{33,78},{90,8},{33,25},{7,50},{49,54},{97,14},{63,14},{99,11},{32,97},{100,20},{74,65},{83,7},{77,79},{92,14},{4,23},{19,38},{9,26},{70,79},
			{35,41},{16,6},{98,23},{58,72},{28,69},{59,30},{71,79},{51,17},{44,85},{22,42},{15,75},{29,1},{41,52},{73,0},{53,98},{23,47},{73,37},{92,50},
			{8,82},{50,50},{3,85},{7,43},{43,36},{59,8},{35,57},{2,84},{37,56},{39,64},{92,14}}, 5))
}