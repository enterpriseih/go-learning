package mid

func removeDuplicateLetters(s string) string {
	// 统计每一个字符出现的次数
	// 单调slice，【循环：如果新的字符比栈顶元素小，而且后面还有栈顶元素，就把栈顶元素丢弃；否则push进去新元素，跳出循环】
	nums := map[byte]int{}
	enterStack := map[byte]int{}
	for i := 0; i < len(s); i++ {
		_, ok := nums[s[i]]
		if !ok {
			nums[s[i]] = 0
			enterStack[s[i]] = 0
		}
		nums[s[i]]++
	}

	stack := []byte{}
	for i := 0; i < len(s); i++ {
		nums[s[i]]--
		for {
			if len(stack) == 0 {
				stack = append(stack, s[i])
				enterStack[s[i]] = 1
				break
			}
			top := stack[len(stack)-1]
			if top > s[i] {
				// 比栈顶的元素小，且没有push到栈里面，未来还有，就可以pop
				if nums[top] > 0 && enterStack[s[i]] == 0 {
					enterStack[stack[len(stack)-1]] = 0
					stack = stack[:len(stack)-1]
					continue
				} else {
					if enterStack[s[i]] == 0 {
						stack = append(stack, s[i])
						enterStack[s[i]] = 1
					}
					break
				}
			} else {
				if enterStack[s[i]] == 0 {
					stack = append(stack, s[i])
					enterStack[s[i]] = 1
				}
				break
			}
		}
	}
	return string(stack)
}
