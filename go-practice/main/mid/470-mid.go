package mid

import "math/rand"

// 7和10的
func Rand10() int {
	for {
		a := rand7()
		b := rand7()
		index := (a-1)*7 + b
		if index <= 40 {
			return (index-1)%10 + 1
		}
	}
}

func rand7() int {
	s2 := rand.NewSource(43)
	r2 := rand.New(s2)
	return r2.Intn(6) + 1
}
