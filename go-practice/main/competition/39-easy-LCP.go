package competition

func minimumSwitchingTimes(source [][]int, target [][]int) int {
	exist := make(map[int]int)
	res := 0

	for i:=0;i<len(target);i++{
		for j:=0;j<len(target[i]);j++ {
			_, ok := exist[target[i][j]]
			if ok {
				exist[target[i][j]]++
			} else {
				exist[target[i][j]] = 1
			}
		}
	}

	for i:=0;i<len(source);i++{
		for j:=0;j<len(source[i]);j++ {
			_, ok := exist[source[i][j]]
			if ok  && exist[source[i][j]] >0 {
				exist[source[i][j]]--
			} else {
				res++
			}
		}
	}
	return res
}