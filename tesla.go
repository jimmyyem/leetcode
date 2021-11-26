package main

import (
	"fmt"
	"sort"
)

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

	// res = nextPermutation([]int{1, 5, 2, 4, 3, 2})

	//res = search([]int{4, 5, 6, 7, 0, 1, 2}, 3)

	params := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	res = groupAnagrams(params)
	fmt.Println("res is", res)
}

// https://leetcode-cn.com/problems/group-anagrams/
// 49. 字母异位词分组
func groupAnagrams(strs []string) [][]string {
	result := [][]string{}
	cache := make(map[string][]string)

	// 1. 遍历每个字符串，并且按字符排序
	// 2. 用map存储每个排序后相同的字符串的集合
	// 3. 遍历map拿到内容存入slice后返回
	for _, str := range strs {
		byteSlice := []byte(str)
		sort.Slice(byteSlice, func(i, j int) bool {
			return byteSlice[i] < byteSlice[j]
		})
		sorted := string(byteSlice)
		if _, ok := cache[sorted]; !ok {
			cache[sorted] = make([]string, 0, 8)
		}
		cache[sorted] = append(cache[sorted], str)
	}
	for _, val := range cache {
		result = append(result, val)
	}

	return result
}

// https://leetcode-cn.com/problems/search-in-rotated-sorted-array/
// 33. 搜索旋转排序数组（二分查找的变种）		4,5,6,7,0,1,2
func search(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}
	low, high := 0, len(nums)-1
	for low <= high {
		mid := low + (high-low)>>1
		if nums[mid] == target {
			return mid
		} else if nums[mid] > nums[low] { // 在数值大的一部分区间里 4,5,6,7 即左边
			if nums[low] <= target && target < nums[mid] { // 如果目标在左边，调整right
				high = mid - 1
			} else { // 不在这部分里，说明在右边，
				low = mid + 1
			}
		} else if nums[mid] < nums[high] { // 在数值小的一部分区间里 0,1,2 即右边
			if nums[mid] < target && target <= nums[high] { // 右边是完全递增的，且target在中间则调整 left
				low = mid + 1
			} else { // 右边不是递增的，调整 right
				high = mid - 1
			}
		} else {
			// 防止出现死循环 一直卡在这里
			if nums[low] == nums[mid] {
				low++
			}
			if nums[high] == nums[mid] {
				high--
			}
		}
	}
	return -1
}

// https://leetcode-cn.com/problems/next-permutation/
// 31. 下一个排列
func nextPermutation(nums []int) []int {
	size := len(nums)
	idx := size - 2
	// 1. 找到反转点，就是找到要交换的小数。 1, 5, 2, 4, 3, 2 这里就是找到 2
	// 找到4后，然后找到后面比4大的交换位置 才能找到比当前大，但是幅度最小的
	for idx >= 0 && nums[idx] >= nums[idx+1] {
		idx--
	}

	// 如果i==-1则说明原始数字是最大的，比如 654321，会跳过下一步搜索123465后比4大的数字过程，而直接进入排序步骤3

	// 2.从后往前找到比2大的第一个数，然后交换。1, 5, 2, 4, 3, 2 j这里最后停留在3这里
	// 交换2和3后是 1, 5, 3, 4, 2, 2
	if idx >= 0 {
		switcher := size - 1
		for switcher >= 0 && nums[idx] >= nums[switcher] {
			switcher--
		}

		nums[idx], nums[switcher] = nums[switcher], nums[idx]
	}

	// 3. 把 3 后面的数字重新从小到大排序，即是最小的比原始数字大的符合条件数字
	// 因为 1, 5, 3, 4, 2, 2 还是比 1, 5, 3, 2, 2, 4 要大，所以 3 后面数字需要重新排序
	reverseSlice(nums[idx+1:])

	return nums
}

func reverseSlice(nums []int) {
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
