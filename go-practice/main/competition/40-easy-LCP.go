package competition

import "sort"

func maxmiumScore(cards []int, cnt int) int {
	odds := []int{}
	evens := []int{}
	sort.Sort(sort.Reverse(sort.IntSlice(cards)))
	for i:=0;i<len(cards);i++{
		if cards[i]%2==0 {
			if len(evens) == 0 {
				evens = append(evens, cards[i])
			} else {
				evens = append(evens, evens[len(evens)-1]+cards[i])
			}
		} else {
			if len(odds) == 0 {
				odds = append(odds, cards[i])
			} else {
				odds = append(odds, odds[len(odds)-1]+cards[i])
			}
		}
	}

	maxRes := 0
	for oddPairCount := 0; oddPairCount*2 <= cnt; oddPairCount++ {
		oneRes := 0
		oddsN := oddPairCount*2
		evensN := cnt-oddPairCount*2
		if oddsN > len(odds) || evensN > len(evens) {
			continue
		}

		if evensN > 0 {
			oneRes = evens[evensN-1]
		}
		if oddsN > 0 {
			oneRes += odds[oddPairCount*2-1]
		}

		if oneRes > maxRes {
			maxRes = oneRes
		}
	}
	return maxRes
}