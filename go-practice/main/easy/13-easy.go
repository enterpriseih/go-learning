package easy

func romanToInt(s string) int {
	convert := make(map[string]int)
	convert["I"] = 1
	convert["V"] = 5
	convert["X"] = 10
	convert["L"] = 50
	convert["C"] = 100
	convert["D"] = 500
	convert["M"] = 1000
	convert["IV"] = 4
	convert["IX"] = 9
	convert["XL"] = 40
	convert["XC"] = 90
	convert["CD"] = 400
	convert["CM"] = 900

	res := 0
	i := 0
	for {
		if i < len(s)-1 {
			v, ok := convert[s[i:i+2]]
			if ok {
				res += v
				i += 2
				// 每次都需要找两个连续的字符
				continue
			}
		}
		if i < len(s) {
			v, ok := convert[s[i:i+1]]
			if ok {
				res += v
				i += 1
			}
		}
		if i >= len(s) {
			break
		}
	}
	return res
}
