package _911

import (
	"container/heap"
	"sort"
)

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func minGroups(intervals [][]int) int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	minRight := &IntHeap{}
	heap.Init(minRight)

	for i := 0; i < len(intervals); i++ {
		if minRight.Len() == 0 || intervals[i][0] <= (*minRight)[0] {
			// open a new group
			heap.Push(minRight, intervals[i][1])
		} else {
			heap.Pop(minRight)
			heap.Push(minRight, intervals[i][1])
		}
	}
	return minRight.Len()
}

func TestMinGroups() int {
	return minGroups([][]int{{5, 10}, {6, 8}, {1, 5}, {2, 3}, {1, 10}})
}
