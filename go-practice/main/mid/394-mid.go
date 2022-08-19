package mid

import (
	"strconv"
)

// 使用栈 394
func decodeString(s string) string {
	// 使用栈，遇到第一个右中括号，就向左找第一个左中括号，然后看数字是多少，把这段数字翻倍，追加到原始字符串里面，再去往后找
	for i := 0; i < len(s); i++ {
		if s[i] == ']' {
			for j := i - 1; j >= 0; j-- {
				if s[j] == '[' {
					repeatStr := s[j+1 : i]
					// repeatNum不一定是一位数，可能是多位数
					nums := ""
					for n := j - 1; n >= 0; n-- {
						if s[n] >= '0' && s[n] <= '9' {
							nums = s[n:n+1] + nums
						} else {
							// 不是数字就要退出的
							break
						}
					}
					println("nums ", nums)
					repeatNum, _ := strconv.Atoi(nums)
					res := ""
					for k := 0; k < repeatNum; k++ {
						res = res + repeatStr
					}
					// pre不能包含数字
					pre := s[0 : j-len(nums)]
					post := s[i+1:]
					s = pre + res + post
					i = j - len(nums) - 1 + repeatNum*len(repeatStr)
					break
				}
			}
		}
	}
	return s
}
