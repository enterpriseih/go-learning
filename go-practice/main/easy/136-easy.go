package easy

func singleNumber(nums []int) int {
	mark := map[int]int{}
	for i := 0; i < len(nums); i++ {
		_, ok := mark[nums[i]]
		if ok {
			delete(mark, nums[i])
		} else {
			mark[nums[i]] = 1
		}
	}

	for k, _ := range mark {
		return k
	}
	return 0
}
