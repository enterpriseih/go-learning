package _911

func mostFrequentEven(nums []int) int {
	counts := map[int]int{}
	for i := 0; i < len(nums); i++ {
		if nums[i]%2 == 0 {
			_, ok := counts[nums[i]]
			if ok {
				counts[nums[i]]++
			} else {
				counts[nums[i]] = 1
			}
		}
	}

	maxCount := -1
	minNum := -1
	for k, v := range counts {
		if v > maxCount {
			maxCount = v
			minNum = k
		} else if v == maxCount {
			if k < minNum || minNum == -1 {
				minNum = k
			}
		}
	}
	return minNum
}

func TestMostFrequentEven() int {
	return mostFrequentEven([]int{29, 47, 21, 41, 13, 37, 25, 7})
}
