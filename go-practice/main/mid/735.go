package mid

import "fmt"

func asteroidCollision(asteroids []int) []int {
	// 只需要考虑两个行星运动距离相反
	// 相邻的一正一负就要合并了
	for {
		flag := false
		for i := 0; i < len(asteroids)-1; i++ {
			if asteroids[i] > 0 && asteroids[i+1] < 0 {
				flag = true
				if -asteroids[i+1] == asteroids[i] {
					if i < len(asteroids)-2 {
						asteroids = append(asteroids[0:i], asteroids[i+2:]...)
					} else {
						asteroids = asteroids[0:i]
					}
					break
				}
				v := 0
				if -asteroids[i+1] > asteroids[i] {
					v = asteroids[i+1]
				} else {
					v = asteroids[i]
				}
				if i < len(asteroids)-2 {
					asteroids = append(asteroids[0:i], append([]int{v}, asteroids[i+2:]...)...)
				} else {
					asteroids = append(asteroids[0:i], v)
				}
				break
			}
		}
		if !flag {
			break
		}
	}
	return asteroids
}

func TestAsteroidCollision() {
	//res := asteroidCollision([]int{5, 10, -5})
	res := asteroidCollision([]int{10, 2, -5})
	println(fmt.Sprintf("%+v", res))
}
