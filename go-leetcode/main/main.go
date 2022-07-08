package main

func main() {
	//PrintHelloWorld()
	//println(minWindow("ADOBECODEBANC", "ABC"))
	//RestoreIpAddresses("101023")
	//var nums []int
	//nums = append(nums, 1)
	//nums = append(nums, 2)
	//nums = append(nums, 3)
	//nums = append(nums, 4)
	//nums = append(nums, 5)
	//fmt.Print("+v", Subsets(nums))
	//println(Rand10())
	a := ListNode{0, nil}
	b := ListNode{4, &a}
	c := ListNode{3, &b}
	d := ListNode{5, &c}
	e := ListNode{-1, &d}
	head := SortList(&e)
	for {
		if head.Next != nil {
			println(head.Val)
			head = head.Next
		} else {
			println(head.Val)
			break
		}
	}
}
