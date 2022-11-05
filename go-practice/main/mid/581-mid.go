package mid

func findUnsortedSubarray(nums []int) int {
	//println("start------")
	startKey := -100001
	startIndex := -1
	endIndex := -1
	endKey := -100001
	maxKey := -100001
	for i := 0; i < len(nums); i++ {
		if maxKey == -100001 || maxKey < nums[i] {
			maxKey = nums[i]
		}
		if nums[i] < maxKey {
			if startIndex == -1 {
				for j := i - 1; j >= 0; j-- {
					if nums[j] <= nums[i] {
						startIndex = j + 1
						startKey = nums[j]
						break
					}
				}
				if startIndex == -1 {
					startIndex = 0
					startKey = nums[i]
				}
			}
			if startIndex != -1 && nums[i] < startKey {
				find := false
				for j := startIndex - 1; j >= 0; j-- {
					if nums[j] <= nums[i] {
						startIndex = j + 1
						startKey = nums[j]
						find = true
						break
					}
				}
				if !find {
					startKey = nums[0]
					startIndex = 0
				}
			}
			endIndex = i
			endKey = nums[i]
		}
		if nums[i] == endKey {
			endIndex = i
		}
		//println(startIndex, endIndex, startKey)
	}
	if endIndex == -1 {
		return 0
	}
	if startIndex == -1 {
		startIndex = 0
	}
	return endIndex - startIndex + 1
}

func TestFindUnsortedSubarray() {
	println(findUnsortedSubarray([]int{-1, -1, -1, -1}) == 0)
	println(findUnsortedSubarray([]int{2, 3, 5, 4, 1}) == 5)
	println(findUnsortedSubarray([]int{1, 2, 5, 3, 4}) == 3)
	println(findUnsortedSubarray([]int{3, 2, 3, 2, 4}) == 4)
	println(findUnsortedSubarray([]int{1, 3, 2, 2, 2}) == 4)
	println(findUnsortedSubarray([]int{2, 1}) == 2)
	println(findUnsortedSubarray([]int{2, 6, 4, 8, 10, 9, 15}) == 5)

	println(findUnsortedSubarray([]int{1, 2, 3, 4}) == 0)
	println(findUnsortedSubarray([]int{1}) == 0)
	println(findUnsortedSubarray([]int{1, 2, 3, 3, 2}) == 3)
}
