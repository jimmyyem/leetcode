package main

import (
	"bytes"
	"container/list"
	"fmt"
	"leetcode/answers"
	"math"
	"sort"
	"strconv"
)

var (
	res interface{}
)

type Element struct {
	key, value int
}

func printList(l *list.List) {
	// 正序
	for e := l.Front(); e != nil; e = e.Next() {
		println(e.Value.(Element).key, e.Value.(Element).value)
	}
	// 倒叙
	//for e := l.Back(); e != nil; e = e.Prev() {
	//	println(e.Value.(Element).key, e.Value.(Element).value)
	//}
	println("==========")
}

func main() {
	//const sample = "\xbd\xb2\x3d\xbc\x20\xe2\x8c\x98"
	//fmt.Println(sample)
	//fmt.Printf("%+q\n", sample)

	//lst := list.New()
	//node1 := Element{
	//	key:   1,
	//	value: 1,
	//}
	//node2 := Element{
	//	key:   2,
	//	value: 2,
	//}
	//node3 := Element{
	//	key:   3,
	//	value: 3,
	//}
	//node4 := Element{
	//	key:   4,
	//	value: 4,
	//}
	//
	//lst.PushFront(node1)
	//printList(lst)
	//
	//lst.PushFront(node2)
	//printList(lst)
	//
	//lst.PushFront(node3)
	//lst.PushBack(node4)
	//printList(lst)

	//res := isValidSymbol("()[]{}")

	//res := twoSum([]int{1, 2, 6, 9, 32}, 10)

	//res = lengthOfLongestSubstring("pwwkew")

	//res = letterCombinations("23")

	// res = nextPermutation([]int{1, 5, 2, 4, 3, 2})

	//res = search([]int{4, 5, 6, 7, 0, 1, 2}, 3)

	//params := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	//res = groupAnagrams(params)

	//res = reverseNumber(123)

	//res = findRepeatedDnaSequences("AAAAAAAAAAA")

	//res = searchRange([]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 3, 4, 4, 4, 4, 5, 5, 6, 6, 6, 8, 10, 10}, 4)

	//res = findClosestElements([]int{1, 2, 3, 4, 5}, 4, 3)

	//res = maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4})

	//res = pivotIndex([]int{1, 7, 3, 6, 5, 6})

	//params := []byte{'a', 'a', 'b', 'b', 'b', 'c', 'c'}
	//res = compress(params)

	//params := []int{1, 2, 2, 3, 4, 5, 6}
	//res = removeElement(params, 2)

	res = missingNum([]int{1, 3, 6, 4, 1, 2})
	fmt.Println("res is", res)
}

// https://leetcode-cn.com/problems/string-compression/
// 443. 压缩字符串，返回压缩后的长度
func compress(chars []byte) int {
	res := []byte{}
	slow, counter := 0, 0
	for slow+counter < len(chars) {
		if chars[slow] == chars[slow+counter] {
			counter++
		} else {
			res = append(res, chars[slow], byte(counter)+'0')
			//fmt.Printf("%v\n", res)
			slow += counter
			counter = 1
		}
	}
	if slow < len(chars) {
		res = append(res, chars[slow], byte(counter)+'0')
	}
	//fmt.Printf("%v\n", res)

	return len(res)
}

// https://leetcode-cn.com/problems/maximum-subarray/
// 53. 最大子数组和
func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	var (
		currValue int
		lastValue = nums[0]
		maxSum    = nums[0]
	)

	for i := 1; i < len(nums); i++ {
		currValue = max(lastValue+nums[i], nums[i])
		if currValue > maxSum {
			maxSum = currValue
		}
		lastValue = currValue
	}

	return maxSum
}

// https://leetcode-cn.com/problems/find-k-closest-elements/
// 658. 找到 K 个最接近的元素
func findClosestElements(arr []int, k int, x int) []int {
	start := 0
	end := len(arr) - k

	for start < end {
		mid := start + (end-start)/2

		// 以一段为移动步长，同时利用二分性质比较最左边与x的差，和最右边与x的差，判断是左移（修改right）还是右移（修改left）
		if x-arr[mid] > arr[mid+k]-x {
			start = mid + 1
		} else {
			end = mid
		}
	}
	//println(start, end, k)

	return arr[start : start+k]
}

// https://leetcode-cn.com/problems/repeated-dna-sequences/
// 187. 重复的DNA序列
func findRepeatedDnaSequences(s string) []string {
	cache := map[string]int{}
	n := len(s)
	res := []string{}

	for i := 0; i+10 <= n; i++ {
		key := s[i : i+10]
		cache[key]++
	}
	fmt.Printf("%+v", cache)

	for str, val := range cache {
		if val > 1 {
			res = append(res, str)
		}
	}

	return res
}

// https://leetcode-cn.com/problems/evaluate-reverse-polish-notation/
// 150. 逆波兰表达式求值
func evalRPN(tokens []string) int {
	buf := make([]int, 0, len(tokens))
	for _, val := range tokens {
		switch val {
		case "+", "-", "*", "/":
			// 区分 + - * / 拿出操作数 做运算
			n := len(buf)
			first := buf[n-2]
			second := buf[n-1]
			buf = buf[:n-2]
			var temp int
			if val == "+" {
				temp = first + second
			} else if val == "-" {
				temp = first - second
			} else if val == "*" {
				temp = first * second
			} else {
				temp = first / second
			}

			buf = append(buf, temp)
		default:
			num, _ := strconv.Atoi(val)
			buf = append(buf, num)
		}
	}

	if len(buf) == 1 {
		return buf[0]
	}

	return 0
}

// https://leetcode-cn.com/problems/lru-cache/
// 146. LRU 缓存机制
type entry struct {
	key, value int
}

type LRUCache struct {
	cap   int
	cache map[int]*list.Element
	lst   *list.List
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		cap:   0,
		cache: map[int]*list.Element{},
		lst:   list.New(),
	}
}

func (c *LRUCache) Get(key int) int {
	value := c.cache[key]
	if value == nil {
		return -1
	}

	c.lst.MoveToFront(value)
	return value.Value.(*entry).value
}

func (c *LRUCache) Put(key, value int) {
	e := c.cache[key]
	if e != nil {
		e.Value = entry{
			key:   key,
			value: value,
		}
		c.lst.MoveToFront(e)
		return
	}

	newElement := c.lst.PushFront(entry{
		key:   key,
		value: value,
	})
	c.cache[key] = newElement

	if len(c.cache) > c.cap {
		lastKey := c.lst.Back().Value.(entry).key
		delete(c.cache, lastKey)
	}

	return
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */

// https://leetcode-cn.com/problems/reorder-list/
// 143. 重排链表
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
//func reorderList(head *ListNode) {
//	nodeList := []*ListNode{}
//	for node := head; node != nil; node = node.Next {
//		nodeList = append(nodeList, node)
//	}
//
//	start, end := 0, len(nodeList)-1
//	for start < end {
//		// 0=>n, 1=>n-1
//		nodeList[start].Next = nodeList[end]
//		start++
//
//		// start++ end-- 时候node可能重叠
//		if start == end {
//			break
//		}
//
//		// n=>1, n-1=>2
//		nodeList[end].Next = nodeList[start]
//		end--
//	}
//	nodeList[start].Next = nil
//	return
//}

// https://leetcode-cn.com/problems/reverse-integer/
// 7. 整数反转
func reverseNumber(x int) int {
	var flag, res = 1, 0
	if x < 0 {
		flag = -1
		x = -x
	}

	for x > 0 {
		res = res*10 + x%10
		x /= 10

		if res > math.MaxInt32 {
			return 0
		}
		if res < math.MinInt32 {
			return 0
		}
	}

	return res * flag
}

// https://leetcode-cn.com/problems/binary-tree-level-order-traversal/
// 102. 二叉树的层序遍历
//func levelOrder(root *TreeNode) (res [][]int) {
//	q := []*TreeNode{root}
//
//	for len(q) > 0 {
//		tmpValue := []int{}
//		tmpNode := []*TreeNode{}
//
//		for _, val := range q {
//			tmpValue = append(tmpValue, val.Value)
//			tmpNode = append(tmpNode, val.Left, val.Right)
//		}
//
//		if len(tmpNode) > 0 {
//			res = append(res, tmpValue)
//			q = tmpNode
//		} else {
//			break
//		}
//	}
//
//	return
//}

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

// https://leetcode-cn.com/problems/find-first-and-last-position-of-element-in-sorted-array/
// 34. 在排序数组中查找元素的第一个和最后一个位置
func searchRange(nums []int, target int) []int {
	getPos := func(nums []int, target int) int {
		var start, end = 0, len(nums) - 1
		for start <= end {
			mid := start + (end-start)/2
			if target == nums[mid] {
				return mid
			} else if target == nums[start] {
				return start
			} else if target == nums[end] {
				return end
			} else if target > nums[start] && target < nums[mid] {
				end = mid - 1
			} else {
				start = mid + 1
			}
		}
		return -1
	}

	pos := getPos(nums, target)
	if pos == -1 {
		return []int{-1, -1}
	}

	// 找到内容后向左、向右延伸
	left, right := pos, pos
	for left-1 >= 0 && nums[left-1] == target {
		left--
	}
	for right+1 < len(nums) && nums[right+1] == target {
		right++
	}

	return []int{left, right}
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
	stack := NewStackTesla()
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

func NewStackTesla() *StackTesla {
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

// https://leetcode-cn.com/problems/find-pivot-index/
// 724. 寻找数组的中心下标	1, 7, 3, 6, 5, 6
func pivotIndex(nums []int) int {
	var leftSum, rightSum, totalSum int
	for _, val := range nums {
		totalSum += val
	}
	// leftSum + nums[i] + rightSum = totalSum
	for i := 0; i < len(nums); i++ {
		rightSum = totalSum - leftSum - nums[i]
		if leftSum == rightSum {
			return i
		}
		leftSum += nums[i]
	}

	return -1
}

// https://mp.weixin.qq.com/s?__biz=Mzg3NDU4MDQ3Mw==&mid=2247491440&idx=2&sn=2937dfd27da220af83bf263d8559a1b2&chksm=cecfcfe6f9b846f045cca09988f47ced4e026d0b5f64f8ee99154a8d5fd8506eddce1c87b31f&cur_album_id=1839734456706760708&scene=190#rd
// 删除链表里重复元素
func deleteDuplicates(head *answers.ListNode) *answers.ListNode {
	slow, fast := head, head
	root := &answers.ListNode{
		Val:  0,
		Next: nil,
	}
	res := root
	for fast != nil {
		if slow.Val == fast.Val {
			fast = fast.Next
		} else {
			root.Next = fast
			slow = fast
			fast = fast.Next
		}
	}

	return res.Next
}

// https://mp.weixin.qq.com/s?__biz=Mzg3NDU4MDQ3Mw==&mid=2247491428&idx=1&sn=941ed312a7f0efc8e21d01b59ba87be9&chksm=cecfcff2f9b846e4e60c8805c745618154da2218d033eee1b3f16f28ce69d8342345ef4c783f&cur_album_id=1839734456706760708&scene=190#rd
// 删除重复元素
func removeElement2(nums []int, target int) int {
	if len(nums) == 0 {
		return 0
	}

	var pos, counter int
	size := len(nums)
	for ; pos < size; pos++ {
		if nums[pos] == target {
			//1,[2],2,3,4,5,6 依次用后面的值覆盖前面的值
			for j := pos; j+1 < size; j++ {
				nums[j] = nums[j+1]
			}
			pos--  // 保留当前索引，1,2,2,3,4,5 如果没有这一步 这种情况会出问题
			size-- // size减少1，由于数组整体减少了1
		} else {
			counter++
		}
	}
	fmt.Printf("%v\n", nums)

	return counter
}

// 双指针（遇到target就略过，遇到不是target就fast后覆盖前slow）
func removeElement(nums []int, target int) int {
	if len(nums) == 0 {
		return 0
	}

	var slow, fast int
	for fast = 0; fast < len(nums); fast++ {
		if nums[fast] != target {
			nums[slow] = nums[fast]
			slow++
		}
	}
	fmt.Printf("%v\n", nums)

	return slow
}

// https://mp.weixin.qq.com/s?__biz=Mzg3NDU4MDQ3Mw==&mid=2247491442&idx=1&sn=753e9d5fda6e6d4c9117df20dc8d3920&source=41#wechat_redirect
// 输入两个链表，找出它们的第一个公共节点
//func getIntersectionNode(headA, headB *answers.ListNode) int {
//	tempA, tempB := headA, headB
//
//	for tempA != tempB {
//		if tempA != nil {
//			tempA = tempA.Next
//		} else {
//			tempA = tempB
//		}
//
//		if tempB != nil {
//			tempB = tempB.Next
//		} else {
//			tempB = tempA
//		}
//	}
//}

//score=55
func missingNum(nums []int) int {
	sort.Ints(nums)
	size := len(nums)

	for i := 0; i+1 < size; i++ {
		if nums[i+1] == nums[i] {
			continue
		}
		if nums[i+1]-nums[i] != 1 && nums[i] > 0 {
			return nums[i] + 1
		}
	}

	if nums[size-1] > 0 {
		return nums[size-1] + 1
	}

	return 1
}

//score=22
func missingNum2(nums []int) int {
	size, maxVal := len(nums), -1<<32

	cache := make(map[int]int, len(nums))

	for i := 0; i < size; i++ {
		cache[nums[i]]++
		if nums[i] > maxVal {
			maxVal = nums[i]
		}
	}
	//fmt.Printf("%v\n", cache)
	for key, _ := range cache {
		if _, ok := cache[key+1]; !ok && key > 0 && key < maxVal {
			return (key + 1)
		}
	}

	if maxVal < 0 {
		return 1
	}

	return maxVal + 1
}

//score=33
func missingNum3(nums []int) int {
	size, maxVal := len(nums), -1<<32

	cache := new(Bitmap)

	for i := 0; i < size; i++ {
		if nums[i] > maxVal {
			maxVal = nums[i]
		}
		if nums[i] > 0 {
			cache.Add(nums[i])
		}
	}

	for i := 0; i < size; i++ {
		next := nums[i] + 1
		if next >= 0 {
			ok := cache.IsExist(next)
			if !ok && nums[i] != maxVal {
				return nums[i] + 1
			}
		}
	}

	if maxVal < 0 {
		return 1
	}

	return maxVal + 1
}

type Bitmap struct {
	// i*64 and get the value (i is the index of slice)
	modValues []uint64
	length    int
}

func (bitmap *Bitmap) IsExist(num int) bool {
	// get the floor and the position in a uint64 number.
	floor, bit := num/64, uint(num%64)
	return floor < len(bitmap.modValues) && (bitmap.modValues[floor]&(1<<bit)) != 0
}

// add a number and return true if succeeds, false if not (it is in already).
func (bitmap *Bitmap) Add(num int) bool {
	floor, bit := num/64, uint(num%64)
	isExist := false
	for floor >= len(bitmap.modValues) {
		bitmap.modValues = append(bitmap.modValues, 0)
		isExist = true
	}
	// get the position in the uint64
	remainder := uint64(1 << bit)

	// judge if the number is in the bitmap.
	if !isExist && bitmap.modValues[floor]&remainder != 0 {
		return false
	}

	isExist = true

	// the remainder index will be set to 1.
	bitmap.modValues[floor] |= remainder
	bitmap.length++

	return isExist
}

// remove a number and return true if succeeds or false if not (the number is not in bitmap).
func (bitmap *Bitmap) Sub(num int) bool {
	floor, bit := num/64, uint(num%64)
	for floor >= len(bitmap.modValues) {
		return false
	}
	// judge the number in the bitmap, directly return false if not.
	if bitmap.modValues[floor]&(1<<bit) == 0 {
		return false
	}

	// the remainder index will be set to 0.
	bitmap.modValues[floor] -= 1 << bit
	bitmap.length--

	return true
}

// the number of the elements in the bitmap.
func (bitmap *Bitmap) Len() int {
	return bitmap.length
}

// output the binary number, from lower to upper.
func (bitmap *Bitmap) BitString() string {
	buf := bytes.Buffer{}
	for _, m := range bitmap.modValues {
		item := strconv.FormatUint(m, 2)
		for i := len(item) - 1; i >= 0; i-- {
			buf.WriteString(item[i : i+1])
		}
	}
	return buf.String()
}
