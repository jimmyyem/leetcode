package answers

//https://leetcode-cn.com/problems/determine-if-string-halves-are-alike/
func HalvesAreAlike(s string) bool {
	var acount ,bcount int

	middle := len(s) >> 1
	for i := 0; i < len(s); i++ {
		if s[i] == 'A' || s[i] == 'a' || s[i] == 'E' || s[i] == 'e' || s[i] == 'I' || s[i] == 'i' || s[i] == 'O' || s[i] == 'o' || s[i] == 'U' || s[i] == 'u' {
			if i < middle {
				acount++
			} else {
				bcount++
			}
		}
	}

	return acount == bcount
}
