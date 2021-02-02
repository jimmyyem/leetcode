package answers

type S917 []string

func (b S917) Len() int {
	return len(b)
}

func (b S917) Less(i, j int) bool {
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

}

//https://leetcode-cn.com/problems/reorder-data-in-log-files/
func ReorderLogFiles(logs []string) []string {

}
