package main

import (
	"fmt"
	"leetcode/answers"
)

func main() {
	//runtime.GOMAXPROCS(8)

	//各种斐波那歇数列
	//start := time.Now()
	//fmt.Println(answers.Fibo2(30))
	//cost := time.Since(start).String()
	//fmt.Println(cost)

	//动态规划，获取最长子序列长度
	//nums := []int{1, 5, 2, 4, 3}
	//cnt := answers.GetAllPath(nums)
	//fmt.Println(cnt)

	//动态规划
	//nums := []int{1, 2, 5}
	//res := answers.CoinChange(nums, 11)
	//fmt.Printf("%v\n", res)

	//快速排序
	//nums := []int{1, 9, 36, 7, 4, 2, 3, 22, 89, 5}
	//answers.QuickSort(nums, 0, len(nums)-1)
	//fmt.Printf("%v\n", nums)

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
	//size := answers.LengthOfLongestSubstring("abcabcbb")
	//fmt.Println(size)

	//???4. 寻找两个正序数组的中位数
	//nums1 := []int{1, 2, 3, 4, 5}
	//nums2 := []int{6}
	//println(answers.FindMedianSortedArrays(nums1, nums2))

	//5.最长回文子串
	//println(answers.LongestPalindrome("babad"))

	//6. Z 字形变换
	//sub := answers.Convert("PAYPALISHIRING", 4)
	//fmt.Println(sub)

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

	//12. 整数转罗马数字
	//res := answers.IntToRoman(3588)
	//fmt.Println(res)

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

	//21. 合并两个有序链表
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

	//42.接雨水
	//nums := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	//println(answers.Trap(nums))

	//46.全排列
	//nums := []int{
	//	1, 2, 3,
	//}
	//res := answers.Permute(nums)
	//dumpSlice2(res)

	//??49. 字母异位词分组
	//strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	//res := answers.GroupAnagrams(strs)
	//fmt.Printf("%v\n", res)

	//动态规划
	//62. 不同路径
	//res := answers.UniquePaths(7, 3)

	//64. 最小路径和
	//grid := make([][]int, 3)
	//for i := 0; i < 3; i++ {
	//	grid[i] = make([]int, 3)
	//}
	//grid[0] = []int{1, 3, 1}
	//grid[1] = []int{1, 5, 1}
	//grid[2] = []int{4, 2, 1}
	//res := answers.MinPathSum(grid)
	//fmt.Println(res)

	//剑指 Offer 61. 扑克牌中的顺子
	//nums := []int{0, 0, 0, 19, 5}
	//flags := answers.IsStraight(nums)
	//println(flags)

	//14.最长共有前缀
	//var strs = []string{}
	//strs = append(strs, "c")
	//strs = append(strs, "acc")
	//strs = append(strs, "ccc")
	//prefix := answers.LongestCommonPrefix(strs)
	//println(prefix)

	//println("计算无符号二进制数里1的个数")
	//var num uint32 = 000000000100
	//println(answers.HammingWeight(num))

	//??121. 买卖股票的最佳时机
	//nums := []int{1, 2, 3, 4, 5}
	//res := answers.MaxProfit121(nums)
	////res := answers.MaxProfit122(nums)
	//fmt.Printf("%v\n", res)

	//150. 逆波兰表达式求值
	//tokens := []string{"4", "13", "5", "/", "+"}
	//res := answers.EvalRPN(tokens)
	//fmt.Println(res)

	//tokens := []string{"9", "+", "(", "3", "-", "1", ")", "*", "3", "+", "10", "/", "2"}
	//res := answers.Infix2Suffix(tokens)
	//fmt.Printf("%v\n", res)

	//190.颠倒二进制数
	//var num uint32 = 4294967293
	//res := answers.ReverseBits(num)
	//println(res)

	//202. 快乐数
	//res := answers.IsHappy(2)
	//println(res)

	//461. 汉明距离
	//res := answers.HammingDistance(1, 4)
	//fmt.Println("distance is ", res)

	//468. 验证IP地址.验证IP是否合法
	//flags := answers.ValidIPAddress("192.0.0.1")
	//println(flags)

	//664. 奇怪的打印机
	//res := answers.StrangePrinter("ababab")
	//fmt.Println(res)

	//226.翻转二叉树
	//993. 二叉树的堂兄弟节点
	//root := &answers.TreeNode{
	//	Val: 1,
	//	Left: &answers.TreeNode{
	//		Val:  2,
	//		Left: nil,
	//		Right: &answers.TreeNode{
	//			Val:   4,
	//			Left:  nil,
	//			Right: nil,
	//		},
	//	},
	//	Right: &answers.TreeNode{
	//		Val:  3,
	//		Left: nil,
	//		Right: &answers.TreeNode{
	//			Val:   5,
	//			Left:  nil,
	//			Right: nil,
	//		},
	//	},
	//}
	//res := answers.InvertTree(root)
	//res := answers.IsCousins(root, 4, 5)
	//fmt.Printf("%v\n", res)

	//937. 重新排列日志文件
	//arr := []string{"a1 9 2 3 1","g1 act car","zo4 4 7","ab1 off key dog","a8 act zoo"}
	//println(arr)
	//res := answers.ReorderLogFiles(arr)
	//dumpSliceString(res)

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

	//剑指 Offer 11. 旋转数组的最小数字
	//var s string = "textbook"
	//res := answers.HalvesAreAlike(s)
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

	//144. 二叉树的前序遍历
	//94. 二叉树的中序遍历
	//145. 二叉树的后序遍历
	//103. 二叉树的锯齿形层序遍历
	//199. 二叉树的右视图
	//root := &answers.TreeNode{
	//	Val:  1,
	//	Left: nil,
	//	Right: &answers.TreeNode{
	//		Val: 2,
	//		Left: &answers.TreeNode{
	//			Val:   3,
	//			Left:  nil,
	//			Right: nil,
	//		},
	//		Right: nil,
	//	},
	//}
	////res := answers.PreorderTraversal(root)
	////res := answers.InorderTraversal(root)
	////res := answers.PostorderTraversal(root)
	//fmt.Printf("%v\n", res)

	//root := &answers.TreeNode{
	//	Val: 1,
	//	Left: &answers.TreeNode{
	//		Val: 3,
	//		Left: &answers.TreeNode{
	//			Val:   5,
	//			Left:  nil,
	//			Right: nil,
	//		},
	//		Right: &answers.TreeNode{
	//			Val:   3,
	//			Left:  nil,
	//			Right: nil,
	//		},
	//	},
	//	Right: &answers.TreeNode{
	//		Val:  2,
	//		Left: nil,
	//		Right: &answers.TreeNode{
	//			Val:   9,
	//			Left:  nil,
	//			Right: nil,
	//		},
	//	},
	//}
	//res := answers.LevelOrder(root)
	//res := answers.LevelOrderBottom(root)
	//res := answers.ZigzagLevelOrder(root)
	//res := answers.RightSideView(root)
	//res := answers.LargestValues(root)
	//res := answers.AverageOfLevels(root)
	//fmt.Printf("%v\n", res)

	//116. 填充每个节点的下一个右侧节点指针
	//root := &answers.Node{
	//	Val: 1,
	//	Left: &answers.Node{
	//		Val: 2,
	//		Left: &answers.Node{
	//			Val:   4,
	//			Left:  nil,
	//			Right: nil,
	//			Next:  nil,
	//		},
	//		Right: &answers.Node{
	//			Val:   5,
	//			Left:  nil,
	//			Right: nil,
	//			Next:  nil,
	//		},
	//		Next: nil,
	//	},
	//	Right: &answers.Node{
	//		Val: 3,
	//		Left: &answers.Node{
	//			Val:   6,
	//			Left:  nil,
	//			Right: nil,
	//			Next:  nil,
	//		},
	//		Right: &answers.Node{
	//			Val:   7,
	//			Left:  nil,
	//			Right: nil,
	//			Next:  nil,
	//		},
	//		Next: nil,
	//	},
	//	Next: nil,
	//}
	//res := answers.Connect(root)
	//fmt.Printf("%v\n", res)

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

	//316. 去除重复字母
	//res := answers.RemoveDuplicateLetters("")
	//fmt.Println(res)

	//692. 前K个高频单词
	//str := []string{"i", "love", "leetcode", "i", "love", "coding"}
	//res := answers.TopKFrequent(str, 2)
	//fmt.Printf("%v\n", res)

	//1716.计算力扣银行的钱
	//res := answers.TotalMoney(10)
	//println(res)

	//1736. 替换隐藏数字得到的最晚时间
	//res := answers.MaximumTime("??:3?")
	//println(res)

	//1710. 卡车上的最大单元数
	//var boxTypes = [][]int{{5, 10}, {2, 5}, {4, 7}, {3, 9}}
	//res := answers.MaximumUnits(boxTypes, 10)
	//println(res)

	//1732. 找到最高海拔
	//var height = []int{-4,-3,-2,-1,4,3,2}
	//res := answers.LargestAltitude(height)
	//println(res)

	//1758. 生成交替二进制字符串的最少操作数
	//var s = "10010100"
	//println(answers.MinOperations(s))

	//滑动窗口
	//arr := []int{1, 2, 3, 4, 5}
	//res := answers.SlidingWindow(2, arr)
	//fmt.Printf("%v\n", res)

	//元素存在就移到最前面，不存在则放在前面
	//arr := []string{"1", "2", "3", "4", "5"}
	//res := answers.MoveToFront("7", arr)
	//fmt.Printf("%v\n", res)

	//剑指 Offer 22. 链表中倒数第k个节点
	//删除有序列表中重复节点
	//head := &answers.ListNode{
	//	Val: 1,
	//	Next: &answers.ListNode{
	//		Val: 2,
	//		Next: &answers.ListNode{
	//			Val: 3,
	//			Next: &answers.ListNode{
	//				Val: 3,
	//				Next: &answers.ListNode{
	//					Val: 3,
	//					Next: &answers.ListNode{
	//						Val:  4,
	//						Next: nil,
	//					},
	//				},
	//			},
	//		},
	//	},
	//}
	//res := answers.GetKthFromEnd(head, 2)
	//res := answers.Middle(head)
	//res := answers.RemoveDupulicateNode(head)

	//head1 := &answers.ListNode{
	//	Val: 1,
	//	Next: &answers.ListNode{
	//		Val: 2,
	//		Next: &answers.ListNode{
	//			Val:  5,
	//			Next: nil,
	//		},
	//	},
	//}
	//head2 := &answers.ListNode{
	//	Val: 1,
	//	Next: &answers.ListNode{
	//		Val: 3,
	//		Next: &answers.ListNode{
	//			Val:  6,
	//			Next: nil,
	//		},
	//	},
	//}
	//res := answers.MergeList(head1, head2)
	//for res != nil {
	//	fmt.Println(res.Val)
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

func dumpTreeList(res *answers.TreeNode) {
	if res != nil {
		fmt.Printf("val is %d\n", res.Val)
		if res.Left != nil {
			dumpTreeList(res.Left)
		}
		if res.Right != nil {
			dumpTreeList(res.Right)
		}
	} else {
		return
	}
}

func max(x, y int) int {
	if x >= y {
		return x
	}
	return y
}

func min(x, y int) int {
	if x >= y {
		return y
	}
	return x
}
