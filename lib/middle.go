package lib

//给定中缀表达式，输出后缀表达式
func InfixToSuffix(s string) string {
	suffix := make([]byte, len(s))
	stack := make([]byte, len(s))

	for i := 0; i < len(s); i++ {
		switch s[i] {
		case '+', '-', '*', '/':

		default:
			if s[i] >= 48 && s[i] <= 57 {
				suffix = append(suffix, s[i])
			}
		}
	}

	return string(suffix)
}

/**
原则：
从左到右遍历中缀表达式的每个数字和符号，若是数字就输出，即成为后缀表达式的一部分；
若是符号，则判断其与栈顶符号的优先级，是右括号或者优先级低于栈顶符号（乘除优先加减）
则栈顶元素依次出栈并输出，并将当前符号进栈，一直到最终输出后缀表达式为止。


9 + ( 3 - 1 ) * 3 + 10 / 2

9 3 1 - 3 * 10 / 2 +
+ +

 */