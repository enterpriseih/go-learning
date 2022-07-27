package main

import "math"

func numSquares(n int) int {
	// 初始化
	// dp[i]= min(dp[i], dp[i-n*n])
	dp := []int{}
	dp = append(dp, 0)
	for i := 1; i <= n; i++ {
		dp = append(dp, i)
		for j := 1; j*j <= i; j++ {
			dp[i] = getMin(dp[i], dp[i-j*j]+1)
		}
	}
	return dp[n]
}

func numSquaresForMath(n int) int {
	// n = (4^a)*(8b+7), return 4
	// n = a*a + b*b , return 0 , 1, 2
	// return 3
	for {
		if n%4 == 0 {
			n = n / 4
		} else {
			break
		}
	}
	if n%8 == 7 {
		return 4
	}
	a := 0
	for {
		b := int(math.Sqrt(float64(n - a*a)))
		if a*a+b*b == n {
			if a == 0 && b == 0 {
				return 0
			}
			if a == 0 || b == 0 {
				return 1
			}
			return 2
		}
		a += 1
		if a*a > n {
			break
		}
	}
	return 3
}
