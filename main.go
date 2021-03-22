package main

import "leetcode/answers"

func main() {
	//fmt.Println("5张扑克牌是否连续")
	//nums := []int{0, 0, 0, 19, 5}
	//flags := answers.IsStraight(nums)
	//fmt.Println(flags)

	//fmt.Println("两数之和")
	//nums := []int{3, 2, 4}
	//target := 6
	//res := answers.TwoSum(nums, target)
	//fmt.Println(res)

	//fmt.Println("验证IP是否合法")
	//flags := answers.ValidIPAddress("192.0.0.1")
	//fmt.Println(flags)

	//fmt.Println("最长共有前缀")
	//var strs []string
	//strs = append(strs, "hahlo")
	//strs = append(strs, "haha")
	//strs = append(strs, "haho")
	//prefix := answers.LongestCommonPrefix(strs)
	//fmt.Println(prefix)

	//fmt.Println("罗马数字转整数")
	//res := answers.RomanToInt("IIV")
	//fmt.Println(res)

	//fmt.Println("计算无符号二进制数里1的个数")
	//var num uint32 = 000000000100
	//fmt.Println(answers.HammingWeight(num))

	//fmt.Println("颠倒二进制数")
	//var num uint32 = 4294967293
	//res := answers.ReverseBits(num)
	//fmt.Println(res)

	//fmt.Println("判断是否是快乐数")
	//res := answers.IsHappy(2)
	//fmt.Println(res)

	//fmt.Println("判断是否是山脉数组")
	//arr := []int{}
	//arr = append(arr,  1,2,3,4)
	//res := answers.ValidMountainArray(arr)
	//fmt.Println(res)

	//fmt.Println("????重新排列日志文件")
	//arr := []string{"a1 9 2 3 1","g1 act car","zo4 4 7","ab1 off key dog","a8 act zoo"}
	//fmt.Println(arr)
	//res := answers.ReorderLogFiles(arr)
	//fmt.Println(res)

	//fmt.Println("左旋转字符串")
	//res := answers.ReverseLeftWords("abcdefg", 2)
	//fmt.Println(res)

	//fmt.Println("第一个只出现一次的字符")
	//res := answers.FirstUniqChar("sasbadxf")
	//fmt.Println(res)

	//fmt.Println("0～n-1中缺失的数字")
	//nums := []int{0, 1, 2, 3, 4, 5, 6, 7, 9}
	//res := answers.MissingNumber(nums)
	//fmt.Println(res)

	//fmt.Println("数组中重复的数字")
	//nums := []int{0, 1, 2, 3, 4, 11, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
	//res := answers.FindRepeatNumber(nums)
	//fmt.Println(res)

	//fmt.Println("旋转数组的最小数字")
	//nums := []int{3, 4, 5, 1, 2}
	//res := answers.MinArray(nums)
	//fmt.Println(res)

	//fmt.Println("旋转数组的最小数字")
	//var s string = "textbook"
	//res := answers.HalvesAreAlike(s)
	//fmt.Println(res)

	//fmt.Println("计算力扣银行的钱")
	//res := answers.TotalMoney(10)
	//fmt.Println(res)

	//fmt.Println("计算力扣银行的钱")
	//res := answers.MaximumTime("??:3?")
	//fmt.Println(res)

	//fmt.Println("卡车上的最大单元数")
	//var boxTypes = [][]int{{5, 10}, {2, 5}, {4, 7}, {3, 9}}
	//res := answers.MaximumUnits(boxTypes, 10)
	//fmt.Println(res)

	//fmt.Println("找到最高海拔")
	//var height = []int{-4,-3,-2,-1,4,3,2}
	//res := answers.LargestAltitude(height)
	//fmt.Println(res)

	//26. 删除有序数组中的重复项
	//var elements = []int{0,0,1,1,1,2,2,3,3,4}
	//res := answers.RemoveDuplicates(elements)
	//println(res)

	//11. 盛最多水的容器
	//var points = []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	//res := answers.MaxArea(points)
	//println(res)

	//15. 三数之和
	//var nums = []int{-2, 0, 3, -1, 4, 0, 3, 4, 1, 1, 1, -3, -5, 4, 0}
	//answers.ThreeSum(nums)

	//704. 二分查找
	//var nums = []int{-1, 0, 3, 5, 9, 12}
	//println(answers.BinarySearch(nums, 9))

	//9. 回文数
	println(answers.IsPalindrome(1234))

	//fmt.Println("合并两个有序链表")
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

	//283.移动零
	//var nums = []int{0, 1, 0, 3, 12}
	//answers.MoveZeroes(nums)

	//1758. 生成交替二进制字符串的最少操作数
	//var s = "10010100"
	//fmt.Println(answers.MinOperations(s))

	//7. 整数反转
	//res := answers.Reverse(-123)
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
	//
	//for res != nil {
	//	println(res.Val, res.Next)
	//	res = res.Next
	//}
}
