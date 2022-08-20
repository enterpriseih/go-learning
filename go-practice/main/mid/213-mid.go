package mid

import "go-practice/main/common"

func rob(nums []int) int {
	// 取第一个或者取最后一个，或者只能取这一个
	if len(nums) == 1 {
		return nums[0]
	}
	return common.GetMax(robLine(nums[0:len(nums)-1]), robLine(nums[1:]))
}

func robLine(nums []int) int {
	// dp[n+1] = max(dp[n-1]+nums[n+1], dp[n])
	var dp []int
	if len(nums) == 0 {
		return 0
	}
	dp = append(dp, nums[0])
	if len(nums) == 1 {
		return dp[0]
	}
	if nums[0] > nums[1] {
		dp = append(dp, nums[0])
	} else {
		dp = append(dp, nums[1])
	}
	for i := 2; i < len(nums); i++ {
		dp = append(dp, common.GetMax(dp[i-2]+nums[i], dp[i-1]))
	}
	return dp[len(nums)-1]
}

func TestRob() {
	println(rob([]int{1, 2, 3}))
}
