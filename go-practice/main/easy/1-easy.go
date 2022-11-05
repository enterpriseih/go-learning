package easy

import (
	"fmt"
)

func twoSum(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		findT := target - nums[i]
		for j := i + 1; j < len(nums); j++ {
			if nums[j] == findT {
				return []int{i, j}
			}
		}

	}
	return []int{}
}

func TestTwoSum() {
	println(fmt.Sprint(twoSum([]int{3, 2, 4}, 6)))
}
