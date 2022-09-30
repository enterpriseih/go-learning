package easy

func repeatedSubstringPattern(s string) bool {
	for i:=1; i<len(s);i++{
		subStr := s[0:i]
		if len(s)%i != 0 {
			continue
		}
		flag := true
		for j:=0;j<len(s);j+=i{
			if j+i <=len(s) && s[j:j+i] !=subStr{
				flag = false
				break
			}
		}
		if flag{
			return true
		}
	}
	return false
}