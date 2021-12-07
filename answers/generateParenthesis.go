package answers

// https://leetcode-cn.com/problems/generate-parentheses/
// 22. 括号生成
func generateParenthesis(n int) []string {

}

//stack结构
type StackPair struct {
	buf  []string
	size int
}

//把一个数据push到stack里
func (s *StackPair) Push(b string) {
	s.buf = append(s.buf, b)
	s.size++
}

//stack里pop出一个数据
func (s *StackPair) Pop() string {
	var result string
	if s.size == 0 {
		return result
	}
	s.size--
	result = s.buf[s.size]
	s.buf = s.buf[0:s.size]

	return result
}

//获取最后一个
func (s *StackPair) GetLast() string {
	var result string
	if s.size == 0 {
		return result
	}

	result = s.buf[s.size-1]
	return result
}
