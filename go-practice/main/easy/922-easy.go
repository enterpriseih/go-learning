package easy

func sortArrayByParityII(nums []int) []int {
	start := 0
	for {
		odd := -1
		even := -1
		for i := start; i < len(nums); i++ {
			if i%2 != nums[i]%2 {
				if i%2 == 0 {
					even = i
				} else {
					odd = i
				}
				start = i
				break
			}
		}

		if odd == -1 && even == -1 {
			return nums
		}

		// find other change index
		if odd == -1 {
			for i := start; i < len(nums); i++ {
				if i%2 == 1 && nums[i]%2 == 0 {
					odd = i
					break
				}
			}
		} else {
			for i := start; i < len(nums); i++ {
				if i%2 == 0 && nums[i]%2 == 1 {
					even = i
					break
				}
			}
		}

		temp := nums[odd]
		nums[odd] = nums[even]
		nums[even] = temp
		println(odd)
		println(even)
		println("-------")
	}
}
