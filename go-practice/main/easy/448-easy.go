package easy

func FindDisappearedNumbers(nums []int) []int {
	i := 0
	for i < len(nums) {
		v := nums[i]
		if nums[v-1] == v {
			i++
		} else {
			temp := nums[v-1]
			nums[v-1] = v
			nums[i] = temp
		}
	}
	var res []int
	for j := 0; j < len(nums); j++ {
		if nums[j] != j+1 {
			res = append(res, j+1)
		}
	}
	return res
}
