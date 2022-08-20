package mid

// 跟316类似，如果一个数比前一个数小就pop，从左到右，因为能pop的数量有限
func removeKdigits(num string, k int) string {
	stack := []byte{}
	popNum := 0
	i := 0
	for {
		if popNum == k && i == len(num) {
			break
		}
		if len(stack) == 0 && i < len(num) {
			stack = append(stack, num[i])
			i++
			continue
		}
		if len(stack) > 0 && i < len(num) {
			// 等于先不pop，小于顶一定要pop（顶不能要）
			if num[i] < stack[len(stack)-1] && popNum < k {
				stack = stack[:len(stack)-1]
				popNum++
			} else {
				stack = append(stack, num[i])
				i++
			}
			continue
		}
		if len(stack) > 0 && popNum < k {
			stack = stack[:len(stack)-1]
			popNum++
		}
	}

	res := string(stack)
	for j := 0; j < len(res); j++ {
		if res[j] != '0' {
			return res[j:]
		}
	}
	if len(res) == 0 || res[0] == '0' {
		return "0"
	}
	return res
}

func TestRemoveKdigits() {
	println(removeKdigits("112", 1))
}
