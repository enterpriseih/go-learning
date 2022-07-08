package main

func minWindow(s string, t string) string {
	minLen := 0
	appearTimes := make(map[byte]int)
	lp := 0
	res := ""

	// 统计t里面的字符数量
	for i := 0; i < len(t); i++ {
		appearTimes[t[i]]++
	}
	for rp := 0; rp < len(s) && lp <= rp; rp++ {
		// 右指针扩展进去
		appearTimes[s[rp]]--
		for {
			if cover(appearTimes) {
				if minLen == 0 || minLen > (rp-lp+1) {
					minLen = rp - lp + 1
					res = s[lp : rp+1]
				}
				// 更新左指针
				appearTimes[s[lp]]++
				lp++
			} else {
				break
			}
		}
	}
	return res
}

func cover(appearTimes map[byte]int) bool {
	for _, v := range appearTimes {
		if v > 0 {
			return false
		}
	}
	return true
}
