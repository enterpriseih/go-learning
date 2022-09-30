package main

import "regexp"

func main() {
	//mid.TestSubarraySum()
	//println(_911.TestMostFrequentEven())
	verion := "1.10.1+"
	pattern := "^[0-9]+.[0-9]+.[0-9]+$"
	matched, _ := regexp.MatchString(pattern, verion)
	println(matched)
}
