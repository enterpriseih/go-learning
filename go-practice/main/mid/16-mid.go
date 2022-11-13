package mid

import (
	"sort"
)

// 太经典的好题，先弱化成两数之和，然后利用双指针，排序后逐渐逼近最优解，复杂度是n平方
func threeSumClosest(nums []int, target int) int {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	//fmt.Printf("%+v", nums)
	closest := nums[0] + nums[1] + nums[2]
	for i:=0; i<len(nums)-2; i++{
		l:=i+1
		r:=len(nums)-1
		for {
			if r <= l{
				break
			}
			cur := nums[i] + nums[l] + nums[r]
			if cur == target {
				return target
			}
			if getAbs(target-closest)  > getAbs(target-cur) {
				closest = cur
			}
			if cur < target{
				l++
			} else {
				r--
			}
		}
	}
	return closest
}

func getAbs(a int) int{
	if a < 0 {
		return -a
	}
	return a
}

func TestThreeSumClosest() {
	print(threeSumClosest([]int{-1,2,1,-4}, 2))
}