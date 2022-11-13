package mid

import "math"

func divide(dividend int, divisor int) int {
	if dividend == math.MinInt32 && divisor == -1 {
		return math.MaxInt32
	}
	positive := true
	if divisor < 0 {
		positive =  !positive
		divisor = -divisor
	}
	if dividend < 0 {
		// dividend、 divisor=math.MinInt32，取反就溢出了
		dividend= -dividend
		positive= !positive
	}

	mark := make(map[int]int)
	for i:=0;; i++ {
		mark[i] = divisor << i
		if mark[i] > dividend {
			break
		}
	}

	nums := len(mark)
	res := 0
	for  {
		nums--
		if nums < 0 || dividend < 0{
			break
		}
		if mark[nums] <= dividend {
			dividend=  dividend - mark[nums]
			res+=1<<nums
		}
	}

	if positive == false{
		return -res
	}
	return res
}

func TestDivide() {
	println(divide(7, -3))
	println(divide(1, 1))
	println(divide(-1, 1))
	println(divide(-2147483648, -1))
}