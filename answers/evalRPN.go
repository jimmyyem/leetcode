package answers

import (
	"strconv"
)

//https://leetcode-cn.com/problems/evaluate-reverse-polish-notation/
//150. 逆波兰表达式求值
func EvalRPN(tokens []string) int {
	if len(tokens) == 0 {
		return 0
	}

	res := 0
	buf := make([]int, 0, len(tokens))

	for i := 0; i < len(tokens); i++ {
		switch tokens[i] {
		case "+", "-", "*", "/":
			size := len(buf)
			//fmt.Printf("%v %d\n", buf, size)
			first := buf[size-2]
			second := buf[size-1]
			buf = buf[:size-2]
			if tokens[i] == "+" {
				buf = append(buf, first+second)
			} else if tokens[i] == "-" {
				buf = append(buf, first-second)
			} else if tokens[i] == "*" {
				buf = append(buf, first*second)
			} else if tokens[i] == "/" {
				if second == 0 {
					buf = append(buf, 0)
				} else {
					buf = append(buf, first/second)
				}
			}
		default:
			num, _ := strconv.Atoi(tokens[i])
			buf = append(buf, num)
		}
	}

	if len(buf) == 1 {
		res = buf[0]
	}

	return res
}

//栈，中缀表达式转为后缀表达式
func Infix2Suffix(tokens []string) []string {
	res := make([]string, 0, len(tokens))
	stack := make([]string, 0, len(tokens))

	for i := 0; i < len(tokens); i++ {
		switch tokens[i] {
		case "+", "-", "*", "/", "(", ")": //符号
			//stack 为空
			if len(stack) == 0 {
				stack = append(stack, tokens[i])
			} else {
				size := len(stack)
				last := stack[size-1]

				//stack不为空
				if tokens[i] == ")" { //从栈顶一直取元素放到res里，遇到（停止
					for j := size - 1; j >= 0; j-- {
						if stack[j] == "(" {
							stack = stack[:j] //栈缩容
							break
						} else {
							res = append(res, stack[j])
						}
					}
				} else if tokens[i] == "(" { //（直接入栈
					stack = append(stack, tokens[i])
				} else {
					//如果当前符号优先级低，则取出栈里所有符号到res，当前符号入栈
					//如果当前符号优先级高，则入栈
					if (last == "*" || last == "/") && (tokens[i] == "+" || tokens[i] == "-") {
						for j := size - 1; j >= 0; j-- {
							res = append(res, stack[j])
						}
						stack = stack[:1]
						stack[0] = tokens[i]
					} else {
						stack = append(stack, tokens[i])
					}
				}
			}
		default: //数字
			res = append(res, tokens[i])
		}
	}

	for k := len(stack) - 1; k >= 0; k-- {
		res = append(res, stack[k])
	}

	return res
}
