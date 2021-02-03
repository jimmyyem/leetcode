package answers

import "sort"

type S937 []string

func (b S937) Len() int {
	return len(b)
}

func (b S937) Less(i, j int) bool {
	m, n := 0, 0

	for ; m < len(b); m++ {
		if b[i][m] == ' ' {
			break
		}
	}
	for ; n < len(b); n++ {
		if b[i][n] == ' ' {
			break
		}
	}

	for m < len(b[i]) && n < len(b[i]) {
		if b[i][m] < b[i][n] {
			return true
		} else if b[i][m] > b[i][n] {
			return false
		} else {
			m++
			n++
		}
	}

	return true
}

func (b S937) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
}

//https://leetcode-cn.com/problems/reorder-data-in-log-files/
func ReorderLogFiles(logs []string) []string {
	var numlogs, wordlogs []string
	for _, v := range logs {
		for k := 0; k < len(v); k++ {
			if v[k] == ' ' {
				if v[k+1] >= 'a' {
					wordlogs = append(wordlogs, v)
				} else {
					numlogs = append(numlogs, v)
				}
				break
			}
		}
	}
	sort.Sort(S937(wordlogs))
	res := append(wordlogs, numlogs...)

	return res
}
