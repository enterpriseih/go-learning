package main

// 718
// 暴力是n的三次方，nums1子串一共有n平方种，然后在nums2中找
// 本质就是找两个子串的最长公共子前缀，dp[i][j]表示nums1从i开始，nums2从j开始的最长公共子前缀
// nums1[i] == nums2[j]: dp[i][j] = dp[i+1][j+1]+1; 否则dp[i][j]=0

func findLength(nums1 []int, nums2 []int) int {
	res := 0
	m := len(nums1)
	n := len(nums2)
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}
	// 最外围初始化为0
	for i := 0; i <= m; i++ {
		dp[i][n] = 0
	}
	for j := 0; j <= n; j++ {
		dp[m][j] = 0
	}

	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			if nums1[i] == nums2[j] {
				dp[i][j] = dp[i+1][j+1] + 1
				if dp[i][j] > res {
					res = dp[i][j]
				}
			} else {
				dp[i][j] = 0
			}
		}
	}
	return res
}
