package main

import "fmt"

var (
	res interface{}
)

func main() {
	const sample = "\xbd\xb2\x3d\xbc\x20\xe2\x8c\x98"
	fmt.Println(sample)
	fmt.Printf("%+q\n", sample)
	//res := isValidSymbol("()[]{}")

	//res := twoSum([]int{1, 2, 6, 9, 32}, 10)

	//res = lengthOfLongestSubstring("pwwkew")

	//res = letterCombinations("23")

	res = nextPermutation([]int{1, 3, 2}) //{1,3,2}
	fmt.Println("res is", res)
}

// https://leetcode-cn.com/problems/next-permutation/
// 31. 下一个排列
func nextPermutation(nums []int) []int {
	n := len(nums)
	i := n - 2
	// 1. 找到反转点，就是找到要交换的小数。 123465 这里就是找到4
	// 倒叙查找效率高
	for i >= 0 && nums[i] >= nums[i+1] {
		i--
	}

	//如果i==-1则说明原始数字是最大的，比如 654321，会跳过下一步搜索123465后比4大的数字过程，而直接进入排序步骤3

	//2.从后往前找到比4大的第一个数，然后交换。123465 j这里最后停留在5这里
	//交换后是123564
	if i >= 0 {
		j := n - 1
		for j >= 0 && nums[i] >= nums[j] {
			j--
		}
		nums[i], nums[j] = nums[j], nums[i]
	}
	//3. 把5后面的数字重新从小到大排序，即是最小的比原始数字大的符合条件数字
	//因为 123564 还是比 123546要大，而 123546 是符合条件的
	reverse(nums[i+1:])
}

func reverse(nums []int) {
	var left, right = 0, len(nums) - 1
	for left < right {
		nums[left], nums[right] = nums[right], nums[left]
		left++
		right--
	}
}

// https://leetcode-cn.com/problems/letter-combinations-of-a-phone-number/
// 17. 电话号码的字母组合
var tables = []string{"", "", "abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz"}

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}
	var result = []string{""}

	for i := 0; i < len(digits); i++ {
		var temp = make([]string, 0, 8)
		alphabet := tables[digits[i]-'0'] //digits[i]是byte类型, 减'0'变剩下0~9之间数
		for j := 0; j < len(alphabet); j++ {
			for k := 0; k < len(result); k++ {
				temp = append(temp, result[k]+string(alphabet[j])) //alphabet[j]是byte类型，需要转成字符串类型才可进行连接操作
			}
		}
		result = temp
	}

	return result
}

// https://leetcode-cn.com/problems/longest-substring-without-repeating-characters/
// 3. 无重复字符的最长子串
func lengthOfLongestSubstring(s string) int {
	if len(s) < 2 {
		return len(s)
	}

	var left, maxSize int
	cache := make(map[byte]int, 26) //key是字符，值是上一次出现的位置

	for right := 0; right < len(s); right++ {
		if _, ok := cache[s[right]]; ok {
			left = max(left, cache[s[right]]+1)
		}
		cache[s[right]] = right
		maxSize = max(maxSize, right-left+1)
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
