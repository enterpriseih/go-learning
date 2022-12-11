package mid

import "fmt"

func insert(intervals [][]int, newInterval []int) [][]int {
	// 只加入一个区间
	l := newInterval[0]
	r := newInterval[1]

	beforeLast := -1
	afterFirst := -1

	lv := l
	rv := r
	for i := 0; i < len(intervals); i++ {
		if l > intervals[i][1] {
			if i == len(intervals)-1 {
				return append(intervals, []int{l, r})
			}
			beforeLast = i
			if intervals[i+1][0] < l {
				lv = intervals[i+1][0]
			} else {
				lv = l
			}
		}
		if r < intervals[i][0] {
			if i == 0 {
				return append([][]int{{l, r}}, intervals...)
			}
			afterFirst = i
			if intervals[i-1][1] > r {
				rv = intervals[i-1][1]
			} else {
				rv = r
			}
			break
		}
	}

	println(beforeLast, afterFirst)

	if beforeLast != -1 && afterFirst != -1 {
		return append(intervals[0:beforeLast+1], append([][]int{{lv, rv}}, intervals[afterFirst:]...)...)
	}

	if beforeLast == -1 && afterFirst == -1 {
		if len(intervals) > 0 && intervals[0][0] < lv {
			lv = intervals[0][0]
		}
		if len(intervals) > 0 && intervals[len(intervals)-1][1] > rv {
			rv = intervals[len(intervals)-1][1]
		}
		return [][]int{{lv, rv}}
	}
	if beforeLast == -1 {
		if intervals[0][0] < lv {
			lv = intervals[0][0]
		}
		return append([][]int{{lv, rv}}, intervals[afterFirst:]...)
	}
	if afterFirst == -1 {
		if intervals[len(intervals)-1][1] > r {
			rv = intervals[len(intervals)-1][1]
		}
		lv = l
		if beforeLast+1 < len(intervals) && intervals[beforeLast+1][0] < lv {
			lv = intervals[beforeLast+1][0]
		}
		return append(intervals[0:beforeLast+1], []int{lv, rv})
	}

	return [][]int{}
}

func TestInsert() {
	//res := insert([][]int{{1, 3}, {6, 9}}, []int{2, 5})
	//res := insert([][]int{{1, 2}, {3, 5}, {6, 7}, {8, 10}, {12, 16}}, []int{4, 8})
	//res := insert([][]int{}, []int{5, 7})
	//res := insert([][]int{{1, 5}}, []int{2, 3})
	//res := insert([][]int{{1, 5}}, []int{2, 7})
	//res := insert([][]int{{2, 4}, {5, 7}, {8, 10}, {11, 13}}, []int{3, 6})
	//res := insert([][]int{{2, 6}, {7, 9}}, []int{15, 18})
	//res := insert([][]int{{0, 3}, {5, 8}, {9, 11}}, []int{9, 18})
	//res := insert([][]int{{3, 5}, {12, 15}}, []int{6, 6})
	//res := insert([][]int{{0, 0}, {1, 3}, {5, 8}, {9, 14}, {15, 17}, {19, 21}}, []int{10, 18})
	res := insert([][]int{{6, 10}, {13, 16}, {19, 19}, {23, 25}, {34, 39}, {41, 43}, {49, 51}}, []int{27, 27})
	println(fmt.Sprintf("%+v", res))
}
