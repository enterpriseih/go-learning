package _007

func minNumBooths(demand []string) int {

	maxMark := map[byte]int{}
	for i := 0; i < len(demand); i++ {
		mark := map[byte]int{}
		for j := 0; j < len(demand[i]); j++ {
			_, ok := mark[demand[i][j]]
			if ok {
				mark[demand[i][j]]++
			} else {
				mark[demand[i][j]] = 1
			}
		}

		for k, v := range mark {
			_, ok := maxMark[k]
			if ok && maxMark[k] < mark[k] {
				maxMark[k] = v
			}
			if !ok {
				maxMark[k] = v
			}
		}
	}

	res := 0
	for _, v := range maxMark {
		res += v
	}
	return res
}

func TestMinNumBooths() {
	println(minNumBooths([]string{"acd", "bed", "accd"}))
}
