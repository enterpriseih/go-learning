package main

// 224
// 找到第一个右括号，然后找左边的第一个左括号，计算出括号中的值
func calculate(s string) int {
	// 当前值的符号1表示加，-1表示减
	res := 0
	sign := 1
	ops := []int{1}
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			ops = append(ops, sign)
		} else if s[i] == ')' {
			ops = ops[:len(ops)-1]
		} else if s[i] == '+' {
			sign = ops[len(ops)-1]
		} else if s[i] == '-' {
			sign = -ops[len(ops)-1]
		} else if s[i] >= '0' && s[i] <= '9' {
			// 找到连续的数字
			nums := 0
			for ; i < len(s) && s[i] >= '0' && s[i] <= '9'; i++ {
				nums = nums*10 + int(s[i]-'0')
			}
			i--
			res += sign * nums
		}
	}
	return res
}
