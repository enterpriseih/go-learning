package mid

func IntToRoman(num int) string {
	replaceRule := make(map[int]string, 13)
	replaceRule[1] = "I"
	replaceRule[5] = "V"
	replaceRule[10] = "X"
	replaceRule[50] = "L"
	replaceRule[100] = "C"
	replaceRule[500] = "D"
	replaceRule[1000] = "M"
	replaceRule[4] = "IV"
	replaceRule[9] = "IX"
	replaceRule[40] = "XL"
	replaceRule[90] = "XC"
	replaceRule[400] = "CD"
	replaceRule[900] = "CM"

	res := ""
	mem := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	for i := 0; i < len(mem); i++ {
		if num < mem[i] {
			continue
		}
		t := 1
		if mem[i] == 1000 || mem[i] == 100 || mem[i] == 10 || mem[i] == 1 {
			t = num / mem[i]
		}
		for ; t > 0; t-- {
			res += replaceRule[mem[i]]
			num -= mem[i]
		}
	}
	return res
}
