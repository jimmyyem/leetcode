package main

import (
	"fmt"
	"leetcode/answers"
)

func main() {
	//runtime.GOMAXPROCS(8)

	//1. 两数之和
	//nums := []int{3, 2, 4}
	//target := 6
	//res := answers.TwoSum(nums, target)
	//println(res)

	//2. 两数相加
	//list1 := answers.ListNode{
	//	Val: 2,
	//	Next: &answers.ListNode{
	//		Val: 4,
	//		Next: &answers.ListNode{
	//			Val:  3,
	//			Next: nil,
	//		},
	//	},
	//}
	//list2 := answers.ListNode{
	//	Val: 5,
	//	Next: &answers.ListNode{
	//		Val: 6,
	//		Next: &answers.ListNode{
	//			Val:  4,
	//			Next: nil,
	//		},
	//	},
	//}
	//res := answers.AddTwoNumbers(&list1, &list2)
	//dumpList(res)

	//3. 无重复字符的最长子串

	//???4. 寻找两个正序数组的中位数
	//nums1 := []int{1, 2, 3, 4, 5}
	//nums2 := []int{6}
	//println(answers.FindMedianSortedArrays(nums1, nums2))

	//5.最长回文子串
	//println(answers.LongestPalindrome("babad"))

	//7. 整数反转
	//res := answers.Reverse(-123)
	//println(res)

	//8.字符串转整数
	//println(answers.MyAtoi(" +12"))

	//9. 回文数
	//println(answers.IsPalindrome(1234))

	//11. 盛最多水的容器
	//var points = []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	//res := answers.MaxArea(points)
	//println(res)

	//13.罗马数字转整数
	//res := answers.RomanToInt("IIV")
	//println(res)

	//15. 三数之和
	//var nums = []int{-2, 0, 3, -1, 4, 0, 3, 4, 1, 1, 1, -3, -5, 4, 0}
	//answers.ThreeSum(nums)

	//16.最接近的3数之和
	//var nums = []int{-1, 2, 1, -4}
	//println(answers.ThreeSumClosest(nums, 1))

	//17.电话号码的字母组合
	//digits := "23"
	//res := answers.LetterCombinations(digits)
	//dumpSliceString(res)

	//42.接雨水
	nums := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	println(answers.Trap(nums))

	//46.全排列
	//nums := []int{
	//	1, 2, 3,
	//}
	//res := answers.Permute(nums)
	//dumpSlice2(res)

	//println("5张扑克牌是否连续")
	//nums := []int{0, 0, 0, 19, 5}
	//flags := answers.IsStraight(nums)
	//println(flags)

	//println("验证IP是否合法")
	//flags := answers.ValidIPAddress("192.0.0.1")
	//println(flags)

	//println("最长共有前缀")
	//var strs []string
	//strs = append(strs, "hahlo")
	//strs = append(strs, "haha")
	//strs = append(strs, "haho")
	//prefix := answers.LongestCommonPrefix(strs)
	//println(prefix)

	//println("计算无符号二进制数里1的个数")
	//var num uint32 = 000000000100
	//println(answers.HammingWeight(num))

	//println("颠倒二进制数")
	//var num uint32 = 4294967293
	//res := answers.ReverseBits(num)
	//println(res)

	//202. 快乐数
	//res := answers.IsHappy(2)
	//println(res)

	//937. 重新排列日志文件
	arr := []string{"a1 9 2 3 1","g1 act car","zo4 4 7","ab1 off key dog","a8 act zoo"}
	println(arr)
	res := answers.ReorderLogFiles(arr)
	dumpSliceString(res)

	//941. 有效的山脉数组
	//arr := []int{}
	//arr = append(arr,  1,2,3,4)
	//res := answers.ValidMountainArray(arr)
	//println(res)

	//println("左旋转字符串")
	//res := answers.ReverseLeftWords("abcdefg", 2)
	//println(res)

	//println("第一个只出现一次的字符")
	//res := answers.FirstUniqChar("sasbadxf")
	//println(res)

	//println("0～n-1中缺失的数字")
	//nums := []int{0, 1, 2, 3, 4, 5, 6, 7, 9}
	//res := answers.MissingNumber(nums)
	//println(res)

	//剑指 Offer 03. 数组中重复的数字
	//nums := []int{2, 3, 1, 0, 2, 5, 3}
	//res := answers.FindRepeatNumber(nums)
	//println(res)

	//剑指 Offer 05. 替换空格
	//var s = "We are happy."
	//res := answers.ReplaceSpace(s)
	//println(res)

	//11. 旋转数组的最小数字
	//nums := []int{3, 4, 5, 1, 2}
	//res := answers.MinArray(nums)
	//println(res)

	//println("旋转数组的最小数字")
	//var s string = "textbook"
	//res := answers.HalvesAreAlike(s)
	//println(res)

	//println("计算力扣银行的钱")
	//res := answers.TotalMoney(10)
	//println(res)

	//println("计算力扣银行的钱")
	//res := answers.MaximumTime("??:3?")
	//println(res)

	//println("卡车上的最大单元数")
	//var boxTypes = [][]int{{5, 10}, {2, 5}, {4, 7}, {3, 9}}
	//res := answers.MaximumUnits(boxTypes, 10)
	//println(res)

	//println("找到最高海拔")
	//var height = []int{-4,-3,-2,-1,4,3,2}
	//res := answers.LargestAltitude(height)
	//println(res)

	//26. 删除有序数组中的重复项
	//var elements = []int{0,0,1,1,1,2,2,3,3,4}
	//res := answers.RemoveDuplicates(elements)
	//println(res)

	//27. 删除元素
	//var elements = []int{3, 2, 2, 3}
	//res := answers.RemoveElement(elements, 2)
	//println(res)

	//31. 下一个排列
	//nums := []int{1, 2, 3, 4, 6, 5}
	//answers.NextPermutation(nums)
	//dumpSlice(nums)

	//33.搜索旋转排序数组
	//var elements = []int{3, 4, 5, 6, 0, 1, 2}
	//res := answers.Search(elements, 2)
	//println(res)

	//35.搜索插入位置
	//var elements = []int{1, 3, 5, 6}
	//res := answers.SearchInsert(elements, 8)
	//println(res)

	//53.最大子序列
	//nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	//res := answers.MaxSubArray(nums)
	//println(res)

	//56. 合并区间
	//var sections = [][]int{
	//	[]int{1, 4},
	//	[]int{0, 2},
	//	[]int{3, 5},
	//}
	//res := answers.MergeSection(sections)
	//dumpSlice2(res)

	//88.合并两个有序数组
	//nums1 := []int{1, 2, 3, 0, 0, 0}
	//nums2 := []int{2, 5, 6}
	//answers.MergeArray(nums1, 3, nums2, 3)
	//dumpSlice(nums1)

	//151. 翻转字符串里的单词
	//var s = "  hello world  "
	//println(answers.ReverseWords(s))

	//154. 寻找旋转排序数组中的最小值
	//var elements = []int{10, 1, 10, 10, 10}
	//res := answers.FindMin(elements)
	//println(res)

	//273.整数转换英文表示
	//num := 1234567891
	//println(answers.NumberToWords(num))

	//704. 二分查找
	//var nums = []int{-1, 0, 3, 5, 9, 12}
	//println(answers.BinarySearch(nums, 9))

	//283.移动零
	//var nums = []int{0, 1, 0, 3, 12}
	//answers.MoveZeroes(nums)

	//1758. 生成交替二进制字符串的最少操作数
	//var s = "10010100"
	//println(answers.MinOperations(s))

	//println("合并两个有序链表")
	//list1 := answers.ListNode{
	//	Val: 1,
	//	Next: &answers.ListNode{
	//		Val: 2,
	//		Next: &answers.ListNode{
	//			Val:  3,
	//			Next: nil,
	//		},
	//	},
	//}
	//list2 := answers.ListNode{
	//	Val: 1,
	//	Next: &answers.ListNode{
	//		Val: 3,
	//		Next: &answers.ListNode{
	//			Val:  4,
	//			Next: nil,
	//		},
	//	},
	//}
	//res := answers.MergeTwoLists(&list1, &list2)
	//
	//for res != nil {
	//	println(res.Val)
	//	res = res.Next
	//}
}

func dumpSlice2(slice [][]int) {
	for _, val := range slice {
		fmt.Printf("%+v", val)
	}
}

func dumpSlice(slice []int) {
	for idx, val := range slice {
		println(idx, "=>", val)
	}
}

func dumpSliceString(slice []string) {
	for idx, val := range slice {
		println(idx, "=>", val)
	}
}

func dumpList(res *answers.ListNode) {
	for res != nil {
		println(res.Val, res.Next)
		res = res.Next
	}
}
