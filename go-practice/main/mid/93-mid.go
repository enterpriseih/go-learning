package mid

import (
	"fmt"
	"strconv"
)

//有效 IP 地址 正好由四个整数（每个整数位于 0 到 255 之间组成，且不能含有前导 0），整数之间用 '.' 分隔。
//
//例如："0.1.2.201" 和 "192.168.1.1" 是 有效 IP 地址，但是 "0.011.255.245"、"192.168.1.312" 和 "192.168@1.1" 是 无效 IP 地址。
//给定一个只包含数字的字符串 s ，用以表示一个 IP 地址，返回所有可能的有效 IP 地址，这些地址可以通过在 s 中插入'.' 来形成。你 不能重新排序或删除 s 中的任何数字。
//你可以按 任何 顺序返回答案。

var res = []string{}

// dfs的调试一定要注意先用一颗小树，大树会绕死你的hh
func RestoreIpAddresses(s string) []string {
	prefix := []string{}
	findRightIPNum(s, 0, prefix)
	println(fmt.Sprintf("+v", res))
	return res
}

func findRightIPNum(str string, depth int, prefixs []string) {
	if depth == 3 {
		if isRightIPNum(str) {
			for _, prefix := range prefixs {
				res = append(res, prefix+"."+str)
			}
		}
		return
	}
	strlen := len(str)
	for i := 0; i < strlen; i++ {
		if isRightIPNum(str[0 : i+1]) {
			needAdd := "."
			newPrefixs := []string{}
			// 不要使用相同的index，过程会很无厘头
			for j, _ := range prefixs {
				newPrefixs = append(newPrefixs, prefixs[j]+needAdd+str[0:i+1])
			}
			// 这里应该是符合ip地址的前缀
			if len(prefixs) == 0 {
				newPrefixs = append(newPrefixs, str[0:i+1])
			}
			findRightIPNum(str[i+1:strlen], depth+1, newPrefixs)
		} else {
			// 如果前面的str已经不符合标准，则也可以剪枝啦，因为再往后数字更大，或者因为前面是0导致本身不符合
			break
		}
	}

}

func isRightIPNum(str string) bool {
	num, err := strconv.Atoi(str)
	if err != nil {
		return false
	}
	if num < 0 || num > 255 {
		return false
	}
	if num != 0 && str[0] == '0' {
		return false
	}
	if num == 0 && len(str) > 1 {
		return false
	}
	return true
}
