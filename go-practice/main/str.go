package main

import (
	"math/rand"
	"strconv"
	"time"
)

func getStr(end int) string {
	s := ""
	rand.Seed(time.Now().Unix())
	for i := 0; i < end; i++ {
		s += "'" + strconv.Itoa(int(i)) + "'"
		s += ","
	}
	return s
}

func getPoint(end int) string {
	//strs := []string{"a-c", "d-f", "g-i", "j-l", "m-o", "p-r", "s-u", "v-x", "y-z", "0-9"}
	rand.Seed(time.Now().Unix())
	s := ""

	ones := getRand6000(320)
	twos := getRandom(320, 604801, 700000)
	for i := 0; i < 320; i++ {
		s += "[" + strconv.Itoa(ones[i]) + "," + strconv.Itoa(twos[i]) + "]" + ","
	}

	ones = getRand6000(916)
	twos = getRandom(916, 518401, 604800)
	for i := 0; i < 916; i++ {
		s += "[" + strconv.Itoa(ones[i]) + "," + strconv.Itoa(twos[i]) + "]" + ","
	}

	ones = getRand6000(284)
	twos = getRandom(284, 432001, 518400)
	for i := 0; i < 284; i++ {
		s += "[" + strconv.Itoa(ones[i]) + "," + strconv.Itoa(twos[i]) + "]" + ","
	}

	ones = getRand6000(300)
	twos = getRandom(300, 345601, 432000)
	for i := 0; i < 300; i++ {
		s += "[" + strconv.Itoa(ones[i]) + "," + strconv.Itoa(twos[i]) + "]" + ","
	}

	ones = getRand6000(220)
	twos = getRandom(220, 259201, 345600)
	for i := 0; i < 220; i++ {
		s += "[" + strconv.Itoa(ones[i]) + "," + strconv.Itoa(twos[i]) + "]" + ","
	}

	ones = getRand6000(50)
	twos = getRandom(50, 172801, 259200)
	for i := 0; i < 50; i++ {
		s += "[" + strconv.Itoa(ones[i]) + "," + strconv.Itoa(twos[i]) + "]" + ","
	}

	ones = getRand6000(230)
	twos = getRandom(230, 86401, 172800)
	for i := 0; i < 230; i++ {
		s += "[" + strconv.Itoa(ones[i]) + "," + strconv.Itoa(twos[i]) + "]" + ","
	}

	return s
}

func getStrIndex(num int) []int {
	res := []int{}
	for i := 0; i < num; i++ {
		res = append(res, int(rand.Int63n(2)))
	}
	return res
}

func getRand6000(num int) []int {
	res := []int{}
	for i := 0; i < num; i++ {
		res = append(res, int(rand.Int63n(6000)))
	}
	return res
}
func getRandom(num, start, end int) []int {
	res := []int{}
	for i := 0; i < num; i++ {
		for {
			two := int(rand.Int63n(700000))
			if two >= start && two < end {
				res = append(res, two)
				break
			}
		}
	}
	return res

}

func randCreator(l int) string {
	str := "0123456789abcdefghigklmnopqrstuvwxyz"
	//strList := []byte(str)
	//
	//result := []byte{}
	//i := 0
	//
	//r := rand.New(rand.NewSource(time.Now().UnixNano()))
	//for i < l {
	//	new := strList[r.Intn(len(strList))]
	//	result = append(result, new)
	//	i = i + 1
	//}
	//return string(result)
	index := int(rand.Int63n(36))
	return str[index : index+1]

}
