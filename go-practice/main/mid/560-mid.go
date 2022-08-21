package mid

func subarraySum(nums []int, k int) int {
	// 暴力枚举
	res := 0
	for i := 0; i < len(nums); i++ {
		temp := 0
		for j := i; j < len(nums); j++ {
			temp += nums[j]
			if temp == k {
				res++
			}
		}
	}
	return res
}

func TestSubarraySum() {
	println(subarraySum([]int{1, 2, 3}, 3))
}
