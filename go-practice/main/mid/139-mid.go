package mid

func wordBreak(s string, wordDict []string) bool {
	// dp[i] = dp[j] && 是否包含子串[j+1,i]
	mark := make(map[string]bool)
	for i := 0; i < len(wordDict); i++ {
		mark[wordDict[i]] = true
	}
	var dp []bool
	for i := 0; i < len(s); i++ {
		flag := false
		for j := 0; j <= i; j++ {
			_, exist := mark[s[j:i+1]]
			if j == 0 && exist {
				flag = true
				break
			}
			if j > 0 && dp[j-1] && exist {
				flag = true
				break
			}
		}
		dp = append(dp, flag)
	}
	return dp[len(s)-1]
}

func TestWordBreak() {
	println(wordBreak("applepenapple", []string{"apple", "pen"}))
}
