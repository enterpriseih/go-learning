package competition

import (
	"fmt"
	"sort"
)

func bicycleYard(position []int, terrain [][]int, obstacle [][]int) [][]int {
	// 很巧妙的题目，不断地朝着四个方向走，如果一个方向(i,j,speed)已经走过了，那么就不用重复往下了
	mark := [][][]int{}
	for i:=0;i<102;i++{
		two := [][]int{}
		for j:=0;j <102;j++ {
			one := make([]int, 102)
			two=append(two, one)
		}
		mark = append(mark, two)
	}

    n := len(terrain)
	m := len(terrain[0])

	moveX := []int{-1,1,0,0}
	moveY := []int{0,0,-1,1}

	layer := [][]int{}
	layer = append(layer, []int{position[0], position[1], 1})
	res   :=  [][]int{}
	markRes := [][]int{}
	for i:=0;i<102;i++{
		one := make([]int, 102)
		markRes=append(markRes, one)
	}
	markRes[position[0]][position[1]] = 1
	for {
		if len(layer) == 0 {
			break
		}
		newlayer := [][]int{}
		for i:=0;i<len(layer);i++ {
			for j:=0;j<4;j++ {
				curX := layer[i][0]
				curY := layer[i][1]
				curS := layer[i][2]
				newX := curX+moveX[j]
				newY := curY+moveY[j]
				if validPos(newX, newY , n, m){
					speed := curS + terrain[curX][curY]-terrain[newX][newY]-obstacle[newX][newY]
					if speed >= 0 && mark[newX][newY][speed]==0 {
						mark[newX][newY][speed] = 1
						newlayer = append(newlayer, []int{newX, newY, speed})
						if speed == 1 && markRes[newX][newY]==0 {
							res = append(res, []int{newX, newY})
							markRes[newX][newY] = 1
						}
					}
				}
			}
		}
		layer = newlayer
	}

	sort.Slice(res, func(i, j int) bool { a, b := res[i], res[j]; return a[0] < b[0] || a[0] == b[0] && a[1] < b[1] })
	return res
}

func validPos(x, y, n,m int) bool{
	if x >=0 && x <n && y>=0 && y<m{
		return true
	}
	return false
}


func TestBicycleYard() {
	//println(fmt.Sprint(bicycleYard([]int{1,1}, [][]int{{5,0},{0,6}}, [][]int{{0,6},{7,0}})))
	println(fmt.Sprint(bicycleYard([]int{0,0}, [][]int{{0,0},{0,0}}, [][]int{{0,0},{0,0}})))
}