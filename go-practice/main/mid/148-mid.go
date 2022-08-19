package mid

// go需要初始化内存 panic: runtime error: invalid memory address or nil pointer dereference
//[signal SIGSEGV: segmentation violation code=0x1 addr=0x0 pc=0x4ac6b2]

// 还有这个指针链表的节点不能换地址
type ListNode struct {
	Val  int
	Next *ListNode
}

var nums []int

func SortList(head *ListNode) *ListNode {
	node := head
	if node == nil {
		return nil
	}
	for {
		if node.Next != nil {
			nums = append(nums, node.Val)
			node = node.Next
		} else {
			nums = append(nums, node.Val)
			break
		}
	}
	mergeSort(0, len(nums)-1)
	t := head
	i := 0
	for t != nil {
		t.Val = nums[i]
		i++
		t = t.Next
	}
	return head
}

func mergeSort(left int, right int) {
	// [left, right]
	if left >= right {
		return
	}
	mid := (left + right) / 2
	mergeSort(left, mid)
	mergeSort(mid+1, right)
	merge(left, mid, mid+1, right)
}

func merge(start1 int, end1 int, start2 int, end2 int) {
	sortStart := start1
	var sorted []int
	for {
		if start1 > end1 || start2 > end2 {
			break
		}
		if nums[start1] <= nums[start2] {
			sorted = append(sorted, nums[start1])
			start1++
		} else {
			sorted = append(sorted, nums[start2])
			start2++
		}
	}
	if start1 > end1 {
		for ; start2 <= end2; start2++ {
			sorted = append(sorted, nums[start2])
		}
	}
	if start2 > end2 {
		for ; start1 <= end1; start1++ {
			sorted = append(sorted, nums[start1])
		}
	}
	for i, _ := range sorted {
		nums[sortStart+i] = sorted[i]
	}
}
