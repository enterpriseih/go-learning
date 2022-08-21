package mid

func findDuplicates(nums []int) []int {
	mark := make(map[int]bool)
	var res []int
	for i := 0; i < len(nums); i++ {
		_, exist := mark[nums[i]]
		if exist {
			res = append(res, nums[i])
		} else {
			mark[nums[i]] = true
		}
	}
	return res
}
