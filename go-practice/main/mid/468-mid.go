package mid

import (
	"strconv"
	"strings"
)

func validIPAddress(queryIP string) string {
	if judgeIPv4(queryIP) {
		return "IPv4"
	}
	if judgeIPv6(queryIP) {
		return "IPv6"
	}
	return "Neither"
}

func judgeIPv4(queryIP string) bool {
	strArray := strings.Split(queryIP, ".")
	if len(strArray) != 4 {
		return false
	}
	for i := 0; i < len(strArray); i++ {
		if len(strArray[i]) == 0 {
			return false
		}
		if i == 0 && strArray[i][0] == '0' && len(strArray[i]) == 1 {
			return false
		}
		if strArray[i][0] == '0' && len(strArray[i]) > 1 {
			return false
		}
		num, err := strconv.Atoi(strArray[i])
		if err != nil {
			return false
		}
		if num < 0 || num > 255 {
			return false
		}
	}
	return true
}

func judgeIPv6(queryIP string) bool {
	strArray := strings.Split(queryIP, ":")
	if len(strArray) != 8 {
		return false
	}
	for i := 0; i < len(strArray); i++ {
		if len(strArray[i]) < 1 || len(strArray[i]) > 4 {
			return false
		}
		for j := 0; j < len(strArray[i]); j++ {
			if !isValidByte(strArray[i][j]) {
				return false
			}
		}

	}
	return true
}

func isValidByte(c byte) bool {
	if c >= '0' && c <= '9' {
		return true
	}
	if c >= 'a' && c <= 'f' {
		return true
	}
	if c >= 'A' && c <= 'F' {
		return true
	}
	return false
}
