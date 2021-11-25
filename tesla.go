package main

import "fmt"

var (
	res interface{}
)

func main() {
	//res := isValidSymbol("()[]{}")

	//res := twoSum([]int{1, 2, 6, 9, 32}, 10)

	res = lengthOfLongestSubstring("pwwkew")

	fmt.Println("res is", res)
}

// https://leetcode-cn.com/problems/longest-substring-without-repeating-characters/
// 3. 无重复字符的最长子串
func lengthOfLongestSubstring(s string) int {
	if len(s) < 2 {
		return len(s)
	}

	var left, maxSize int
	cache := make(map[byte]int, 26) //key是字符，值是上一次出现的位置

	// 循环的时候实现了右边界的每次右移
	for right := 0; right < len(s); right++ {
		// 如果出现过了让左边界右移
		if _, ok := cache[s[right]]; ok {
			left = max(left, cache[s[right]]+1)
		}
		// 每次经过一个点都要记录位置
		cache[s[right]] = right

		// 每次都比较区间长度是否是最大的
		maxSize = max(maxSize, right-left+1)
		//fmt.Println(left, right)	//(left,right) 这是最大子串区间
	}
	return maxSize
}

// https://leetcode-cn.com/problems/two-sum/
// 1.  2数之和
func twoSum(nums []int, target int) []int {
	var cache = make(map[int]int)

	for index, value := range nums {
		left := target - value
		if pos, ok := cache[left]; ok {
			return []int{index, pos}
		} else {
			cache[value] = index
		}
	}

	return []int{}
}

// https://leetcode-cn.com/problems/valid-parentheses/
// 20. 有效的括号
func isValidSymbol(s string) bool {
	stack := NewStack()
	for _, val := range s {
		//fmt.Printf("%v\n", val)
		switch val {
		case '[', '{', '(':
			stack.Push(val)
		case ']':
			var last = stack.Pop()
			if last != '[' {
				return false
			}
		case '}':
			var last = stack.Pop()
			if last != '{' {
				return false
			}
		case ')':
			var last = stack.Pop()
			if last != '(' {
				return false
			}
		}
	}

	if stack.Size != 0 {
		return false
	}

	return true
}

type StackTesla struct {
	Data []rune
	Size int
}

func (s *StackTesla) Push(num rune) {
	s.Data = append(s.Data, num)
	s.Size++
}

func (s *StackTesla) Pop() rune {
	if s.Size == 0 {
		return rune(0)
	}
	var size = s.Size
	var last = s.Data[size-1]
	s.Data = s.Data[:size-1]
	s.Size--
	return last
}

func NewStack() *StackTesla {
	return &StackTesla{
		Data: make([]rune, 0, 8),
		Size: 0,
	}
}

func max(x, y int) int {
	if x >= y {
		return x
	}
	return y
}
