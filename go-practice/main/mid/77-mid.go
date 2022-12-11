package mid

import "fmt"

func combine(n int, k int) [][]int {
	nums := []int{}
	for i := 0; i < n; i++ {
		nums = append(nums, i+1)
	}

	res := [][]int{}
	path := []int{}
	var dfsCombine func(depth int, start int)
	dfsCombine = func(depth int, start int) {
		if depth == k {
			res = append(res, append([]int{}, path...))
			return
		}
		for i := start; i < n; i++ {
			path = append(path, nums[i])
			dfsCombine(depth+1, i+1)
			path = path[:len(path)-1]
		}

	}
	dfsCombine(0, 0)
	return res

}

func TestCombine() {
	res := combine(4, 2)
	println(fmt.Sprintf("%+v", res))
}
