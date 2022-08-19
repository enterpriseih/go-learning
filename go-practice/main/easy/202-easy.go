package easy

func isHappy(n int) bool {
	mark := map[int]int{}
	for {
		if n == 1 {
			return true
		}
		_, exist := mark[n]
		if exist {
			return false
		}
		mark[n] = 1
		newNum := 0
		for n > 0 {
			newNum += (n % 10) * (n % 10)
			n /= 10
		}
		n = newNum
	}
}
