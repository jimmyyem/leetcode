package answers

//https://leetcode-cn.com/problems/first-unique-character-in-a-string/
//387. 字符串中的第一个唯一字符
func FirstUniqChar(s string) byte {
	if len(s) == 0 {
		return ' '
	}

	var list [26]int

	for i := 0; i < len(s); i++ {
		list[s[i]-'a']++
	}

	for i := 0; i < len(s); i++ {
		temp := s[i] - 'a'
		if list[temp] == 1 {
			return s[i]
		}
	}

	return ' '
}
