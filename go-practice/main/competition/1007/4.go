package _007

func Leetcode(words []string) int {
	// 取13次，每次可能取任何一个字母, 然后取这个字母的时候找最小的取，如果最小的有好几个，那也都要找一遍

}

func takeWord(need map[byte]int, words []string, word_cost int) {
	if len(need) == 0 {
		return
	}
	for k, _ := range need {
		takePos, c_cost := findRes(words, k)
		for i := 0; i < len(takePos); i++ {
			new_words := []string{}
			new_need := map[byte]int{}
			copy(new_words, words)
			new_words[takePos[i][0]] = new_words[takePos[i][0]][0:takePos[i][1]] + new_words[takePos[i][0]][takePos[i][1]+1:]
			new_need[k]--
			if new_need[k] == 0 {
				delete(new_need, k)
			}
			takeWord(new_need, new_words, word_cost+c_cost)
		}
	}

}

func findRes(words []string, target byte) ([][]int, int) {
	cost := -1
	takePos := [][]int{}
	for i := 0; i < len(words); i++ {
		for j := 0; j < len(words[i]); j++ {
			if target != words[i][j] {
				continue
			}
			c_cost := j * (len(words[i]) - 1 - j)
			if c_cost == cost || cost == -1 {
				takePos = append(takePos, []int{i, j})
			} else if c_cost < cost {
				takePos = [][]int{}
				takePos = append(takePos, []int{i, j})
			}

		}
	}
	return takePos, cost
}
