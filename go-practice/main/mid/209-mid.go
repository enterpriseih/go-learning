package mid

// 209
func minSubArrayLen(target int, nums []int) int {
	l := 0
	r := 0
	// [l, r]
	sum := nums[0]
	res := 0
	for {
		if sum < target {
			r++
			if r >= len(nums) {
				break
			}
			sum += nums[r]
		} else {
			if res == 0 || r-l+1 < res {
				res = r - l + 1
			}
			sum -= nums[l]
			l++
			if l >= len(nums) {
				break
			}
		}
	}
	return res
}
