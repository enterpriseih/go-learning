package hard

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	length := len(nums1) + len(nums2)
	chooseTwo := false
	findIndexStart := length / 2
	keys := []int{}

	if length%2 == 0 {
		chooseTwo = true
		findIndexStart = length/2 - 1
	}
	i := 0
	j := 0
	for {
		println(i, j, findIndexStart)
		var choosedKey int
		if i == len(nums1) {
			choosedKey = nums2[j]
			j++
		} else if j == len(nums2) {
			choosedKey = nums1[i]
			i++
		} else if i < len(nums1) && j < len(nums2) {
			if nums1[i] > nums2[j] {
				choosedKey = nums2[j]
				j++
			} else {
				choosedKey = nums1[i]
				i++
			}
		}
		if i+j-1 == findIndexStart {
			keys = append(keys, choosedKey)
			if !chooseTwo {
				break
			}
		}
		if i+j-1 == findIndexStart+1 && chooseTwo {
			keys = append(keys, choosedKey)
			break
		}

	}
	if chooseTwo {
		return float64(keys[0]+keys[1]) / 2.0
	}
	return float64(keys[0])
}

func TestFindMedianSortedArrays() {
	println(findMedianSortedArrays([]int{1, 2}, []int{3, 4}))
}
