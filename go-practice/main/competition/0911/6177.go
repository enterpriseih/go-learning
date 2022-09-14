package _911

func partitionString(s string) int {
	counts := map[byte]int{}
	for i := 0; i < len(s); i++ {
		_, ok := counts[s[i]]
		if ok {
			return partitionString(s[i:]) + 1
		} else {
			counts[s[i]] = 1
		}
	}
	return 1
}

func TestPartitionString() int {
	return partitionString("ssssss")
}
