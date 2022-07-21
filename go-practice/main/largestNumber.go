package main

import (
	"sort"
	"strconv"
)

func largestNumber(nums []int) string {
	strs := []string{}
	for i := 0; i < len(nums); i++ {
		strs = append(strs, strconv.Itoa(nums[i]))
	}
	sort.Slice(strs, func(i, j int) bool {
		return strs[i]+strs[j] > strs[j]+strs[i]
	})
	res := ""
	for i := 0; i < len(strs); i++ {
		res += strs[i]
	}

	if len(res) > 0 && res[0] == '0' {
		return "0"
	}
	return res
}
