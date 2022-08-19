package mid

import "go-practice/main"

func maxProduct(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	minRes := []int{}
	maxRes := []int{}

	minRes = append(minRes, nums[0])
	maxRes = append(maxRes, nums[0])

	res := maxRes[0]

	for i := 1; i < len(nums); i++ {
		values1 := minRes[i-1] * nums[i]
		values2 := maxRes[i-1] * nums[i]
		minRes = append(minRes, main.getMin(nums[i], main.getMin(values1, values2)))
		maxRes = append(maxRes, getMax(nums[i], getMax(values1, values2)))
		res = getMax(res, maxRes[i])
	}
	return res
}

func getMax(x1, x2 int) int {
	if x1 > x2 {
		return x1
	} else {
		return x2
	}
}
