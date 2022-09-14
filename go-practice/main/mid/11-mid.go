package mid

import "go-practice/main/common"

func maxAreaFunc(height []int) int {
	res := 0
	for i := 1; i < len(height); i++ {
		for j := 0; j < i; j++ {
			one := common.GetMin(height[i], height[j]) * (i - j)
			if one > res {
				res = one
			}
		}
	}
	return res
}
