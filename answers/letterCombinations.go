package answers

var table []string = []string{"", "", "abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz"}

//https://leetcode-cn.com/problems/letter-combinations-of-a-phone-number/
//17.电话号码的字母组合
func LetterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}
	result := []string{""}

	for i := 0; i < len(digits); i++ {
		t := table[digits[i]-'0']
		temp := []string{}
		for j := 0; j < len(t); j++ {
			for z := 0; z < len(result); z++ {
				temp = append(temp, result[z]+string(t[j]))
			}
		}
		result = temp
	}

	return result
}
