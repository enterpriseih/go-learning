package easy

func getRow(rowIndex int) []int {
	pre := []int{1}
	curRow := []int{}
	for i := 1; i <= rowIndex; i++ {
		curRow = append(curRow, pre[0])
		for j := 0; j < len(pre)-1; j++ {
			curRow = append(curRow, pre[j]+pre[j+1])
		}
		curRow = append(curRow, pre[len(pre)-1])
		pre = curRow
		curRow = []int{}
	}
	return pre
}
