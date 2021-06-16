package answers

//https://leetcode-cn.com/problems/valid-parentheses/solution/
//20. 有效的括号
func IsValid(s string) bool {
	stack := &Stack{
		buf:  make([]rune, 0),
		size: 0,
	}
	for _, str := range s {
		switch str {
		case '{', '[', '(':
			stack.Push(str)
		case ')':
			last := stack.GetLast()
			if last == '(' {
				stack.Pop()
			} else {
				return false
			}
		case ']':
			last := stack.GetLast()
			if last == '[' {
				stack.Pop()
			} else {
				return false
			}
		case '}':
			last := stack.GetLast()
			if last == '{' {
				stack.Pop()
			} else {
				return false
			}
		default:
			return false
		}
	}

	if stack.size > 0 {
		return false
	}

	return true
}

//stack结构
type Stack struct {
	buf  []rune
	size int
}

//把一个数据push到stack里
func (s *Stack) Push(b rune) {
	s.buf = append(s.buf, b)
	s.size++
}

//stack里pop出一个数据
func (s *Stack) Pop() rune {
	var result rune
	if s.size == 0 {
		return result
	}
	s.size--
	result = s.buf[s.size]
	s.buf = s.buf[0:s.size]

	return result
}

//获取最后一个
func (s *Stack) GetLast() rune {
	var result rune
	if s.size == 0 {
		return result
	}

	result = s.buf[s.size-1]
	return result
}
