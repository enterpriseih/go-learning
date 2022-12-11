package mid

func isInterleave(s1 string, s2 string, s3 string) bool {
	m := len(s1)
	n := len(s2)
	if len(s3) != m+n {
		return false
	}
	dp := [][]int{}
	for i := 0; i <= m; i++ {
		dp = append(dp, make([]int, n+1))
	}

	dp[0][0] = 1
	for j := 1; j <= n; j++ {
		if dp[0][j-1] == 1 && s3[j-1] == s2[j-1] {
			dp[0][j] = 1
		}
	}

	for i := 1; i <= m; i++ {
		if dp[i-1][0] == 1 && s3[i-1] == s1[i-1] {
			dp[i][0] = 1
		}
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			// 可能从s1匹配， 也可能从s2匹配
			if dp[i][j-1] == 1 && s3[i+j-1] == s2[j-1] {
				dp[i][j] = 1
			}
			if dp[i-1][j] == 1 && s3[i+j-1] == s1[i-1] {
				dp[i][j] = 1
			}
		}
	}
	if dp[m][n] == 1 {
		return true
	}

	return false
}

func TestIsInterleave() {
	//res := isInterleave("", "", "")
	//res := isInterleave("aabcc", "dbbca", "aadbbbaccc")
	res := isInterleave("aabcc", "dbbca", "aadbbcbcac")
	println(res)
}
