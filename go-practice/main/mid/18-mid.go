package mid

import (
	"fmt"
	"sort"
)

func fourSum(nums []int, target int) [][]int {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})

	sum := 0
	var res [][]int
	for i:=0; i<len(nums)-3; i++ {
		for j:=i+1; j<len(nums)-2; j++ {
			l:=j+1
			r:=len(nums)-1
			for {
				if r <= l {
					break
				}
				sum = nums[i] + nums[j] + nums[l] + nums[r]
				if sum < target {
					l++
				} else if sum == target {
					oneRes := []int{nums[i], nums[j], nums[l], nums[r]}
					println(fmt.Sprintf("%+v", oneRes))
					println(fmt.Sprintf("%+v", res))
					if !judgeIsIn(oneRes, res){
							res = append(res, oneRes)
					}
					l++

				} else {
					r--
				}
			}
		}
	}
	return res
}

func TestFourSum() {
	//println(fmt.Sprintf("%+v", fourSum([]int{2,2,2,2,2}, 8)))
	println(fmt.Sprintf("%+v", fourSum([]int{-5,-4,-3,-2,-1,0,0,1,2,3,4,5}, 0)))
}

func judgeIsIn(ele []int, array [][]int) bool {
	for i :=len(array)-1; i>=0; i--{
		if array[i][0] == ele[0] && array[i][1] == ele[1] && array[i][2] == ele[2] && array[i][3] == ele[3] {
			return true
		}
	}
	return false
}