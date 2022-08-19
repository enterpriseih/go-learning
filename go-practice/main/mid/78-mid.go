package mid

// [1,2,3]
// [1]
// [1, 2] [1]
func Subsets(nums []int) [][]int {
	return find(nums)
}

// nums增加一个index，可能的解答：增加index之前的答案里面（增加新的index，或者不增）
// oneRes 要重新添加而不是直接赋值otherNum
func find(nums []int) [][]int {
	var res [][]int
	var oneRes []int
	if len(nums) == 0 {
		res = append(res, oneRes)
		return res
	}
	if len(nums) > 0 {
		otherNums := find(nums[1:])
		for i, _ := range otherNums {
			res = append(res, otherNums[i])
		}
		for _, otherNum := range otherNums {
			oneRes = []int{}
			for _, j := range otherNum {
				oneRes = append(oneRes, j)
			}
			oneRes = append(oneRes, nums[0])
			res = append(res, oneRes)
		}
	}
	//fmt.Println("%d, %+v", nums[0], res)
	return res
}
