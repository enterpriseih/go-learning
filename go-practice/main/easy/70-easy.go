package easy

//func climbStairs(n int) int {
//	if n == 1 || n == 0 {
//		return 1
//	}
//	return climbStairs(n-1) + climbStairs(n-2)
//}

func climbStairs(n int) int {
	if n == 1 || n == 0 {
		return 1
	}
	a := 1
	b := 1
	for i := 2; i <= n; i++ {
		c := a + b
		a = b
		b = c
	}
	return b

}
