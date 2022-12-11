package mid

import (
	"fmt"
	"sort"
)

func subsetsWithDup(nums []int) [][]int {
	sort.Ints(nums)

	resInt := [][]int{}

	var subDfs func(int, []int)
	subDfs = func(index int, subset []int) {
		resInt = append(resInt, append([]int{}, subset...))

		for i := index; i < len(nums); i++ {
			if i > index && nums[i] == nums[i-1] {
				continue
			}
			subDfs(i+1, append(subset, nums[i]))
		}
	}
	subDfs(0, []int{})
	return resInt
}

func subsetsWithDup1(nums []int) [][]int {
	sort.Ints(nums)

	resInt := [][]int{}
	subset := []int{}

	var subDfs func(bool, int)
	subDfs = func(choosed bool, index int) {
		if index == len(nums) {
			resInt = append(resInt, append([]int{}, subset...))
			return
		}

		if !(!choosed && index > 0 && nums[index] == nums[index-1]) {
			subset = append(subset, nums[index])
			subDfs(true, index+1)
			subset = subset[0 : len(subset)-1]
		}
		subDfs(false, index+1)
	}
	subDfs(false, 0)
	return resInt
}

func TestSubsetsWithDup() {
	res := subsetsWithDup1([]int{2, 1, 2, 1, 3})
	println(fmt.Sprintf("%+v", res))
}
