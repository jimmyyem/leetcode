package main

import (
	"fmt"
	"leetcode/answers"
	"log"
	"runtime"
	"sort"
	"strconv"
	"strings"
)

func main() {
	// go 自带的 sort.Slice 排序
	//rand.Seed(time.Now().UnixNano())
	//old := []lsort.Student{}
	//for i := 0; i < 10; i++ {
	//	old = append(old, lsort.Student{
	//		ID:    i + 1,
	//		Name:  "ID" + strconv.Itoa(i+1),
	//		Score: rand.Intn(99),
	//	})
	//}
	//lsort.DiySort(old)

	//fmt.Println(countAndSay(5))
	//
	//nums := []int{1, 2, 1, 2, 4, 7, 7}
	//fmt.Println(singleNumber(nums))
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

	//各种排序
	//nums := []int{1, 9, 36, 7, 4, 2, 3, 22, 89, 5}
	//lsort.Merge(nums, 0, len(nums)-1)
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

	//14.最长共有前缀
	//var strs = []string{}
	//strs = append(strs, "c")
	//strs = append(strs, "acc")
	//strs = append(strs, "ccc")
	//prefix := answers.LongestCommonPrefix(strs)
	//println(prefix)

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

	//19. 删除链表的倒数第 N 个结点
	//head := &answers.ListNode{
	//	Val: 1,
	//	Next: &answers.ListNode{
	//		Val: 2,
	//		Next: &answers.ListNode{
	//			Val: 3,
	//			Next: &answers.ListNode{
	//				Val: 4,
	//				Next: &answers.ListNode{
	//					Val:  5,
	//					Next: nil,
	//				},
	//			},
	//		},
	//	},
	//}
	//res := answers.RemoveNthFromEnd(head, 2)
	//fmt.Printf("%v\n", res)
	//dumpList(res)

	//20. 有效的括号
	//isValid := answers.IsValid("()[]{}")
	//fmt.Println(isValid)

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

	//??28. 实现 strStr()
	//idx := answers.KMP("ABCDABD&ABCDABD&", "ABCDABD&")
	//idx := answers.QuerySameFix("ABCDAB")
	//fmt.Println("ABCDAB common fix length is ", idx)

	//34. 在排序数组中查找元素的第一个和最后一个位置
	//srange := answers.SearchRange([]int{5, 7, 7, 8, 8, 10}, 8)
	//fmt.Printf("%v", srange)

	//42.接雨水
	//nums := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	//println(answers.Trap(nums))

	//46.全排列
	//nums := []int{
	//	1, 2, 3,
	//}
	//res := answers.Permute(nums)
	//dumpSlice2(res)

	//48. 旋转图像
	//matrix := [][]int{
	//	{1, 2, 3},
	//	{4, 5, 6},
	//	{7, 8, 9},
	//}
	//answers.Rotate(matrix)
	//fmt.Printf("%v\n", matrix)

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

	//66. 加一
	//res := answers.PlusOne([]int{7, 2, 8, 5, 0, 9, 1, 2, 9, 5, 3, 6, 6, 7, 3, 2, 8, 4, 3, 7, 9, 5, 7, 7, 4, 7, 4, 9, 4, 7, 0, 1, 1, 1, 7, 4, 0, 0, 6})
	//fmt.Println(res)

	//76. 最小覆盖子串
	//res := answers.MinWindow("abc", "ac")
	//fmt.Println("res is", res)

	//98. 验证二叉搜索树
	//root := &answers.TreeNode{
	//	Val: 2,
	//	Left: &answers.TreeNode{
	//		Val:   1,
	//		Left:  nil,
	//		Right: nil,
	//	},
	//	Right: &answers.TreeNode{
	//		Val:   3,
	//		Left:  nil,
	//		Right: nil,
	//	},
	//}
	//res := answers.IsValidBST(root)
	//fmt.Println(res)

	//101. 对称二叉树
	//root := &answers.TreeNode{
	//	Val: 1,
	//	Left: &answers.TreeNode{
	//		Val: 2,
	//		Left: &answers.TreeNode{
	//			Val:   3,
	//			Left:  nil,
	//			Right: nil,
	//		},
	//		Right: &answers.TreeNode{
	//			Val:   4,
	//			Left:  nil,
	//			Right: nil,
	//		},
	//	},
	//	Right: &answers.TreeNode{
	//		Val: 2,
	//		Left: &answers.TreeNode{
	//			Val:   4,
	//			Left:  nil,
	//			Right: nil,
	//		},
	//		Right: &answers.TreeNode{
	//			Val:   3,
	//			Left:  nil,
	//			Right: nil,
	//		},
	//	},
	//}
	//res := answers.IsSymmetric(root)
	//fmt.Println(res)

	//104. 二叉树的最大深度
	//root := &answers.TreeNode{
	//	Val: 1,
	//	Left: nil,
	//	Right: &answers.TreeNode{
	//		Val: 20,
	//		Left: nil,
	//		Right: nil,
	//	},
	//}
	//depth := answers.MaxDepth(root)
	//fmt.Println(depth)

	//剑指 Offer 61. 扑克牌中的顺子
	//nums := []int{0, 0, 0, 19, 5}
	//flags := answers.IsStraight(nums)
	//println(flags)

	//剑指 Offer 15. 二进制中1的个数
	//var num uint32 = 0000011100100
	//println(answers.HammingWeight(num))

	//??121. 买卖股票的最佳时机
	//nums := []int{1, 2, 3, 4, 5}
	//res := answers.MaxProfit121(nums)
	////res := answers.MaxProfit122(nums)
	//fmt.Printf("%v\n", res)

	//141. 环形链表

	//150. 逆波兰表达式求值
	//tokens := []string{"4", "13", "5", "/", "+"}
	//res := answers.EvalRPN(tokens)
	//fmt.Println(res)

	//tokens := []string{"9", "+", "(", "3", "-", "1", ")", "*", "3", "+", "10", "/", "2"}
	//res := answers.Infix2Suffix(tokens)
	//fmt.Printf("%v\n", res)

	//155. 最小栈

	//190.颠倒二进制数
	//var num uint32 = 4294967293
	//res := answers.ReverseBits(num)
	//println(res)

	//198. 打家劫舍
	//maxVal := answers.Rob([]int{2, 1, 1, 2})
	//fmt.Println(maxVal)

	//202. 快乐数
	//res := answers.IsHappy(2)
	//println(res)

	//206. 反转链表
	//234. 回文链表
	//list := &answers.ListNode{
	//	Val: 1,
	//	Next: &answers.ListNode{
	//		Val: 2,
	//		Next: &answers.ListNode{
	//			Val: 2,
	//			Next: &answers.ListNode{
	//				Val:  1,
	//				Next: nil,
	//			},
	//		},
	//	},
	//}
	//list = answers.ReverseList(list)
	//dumpList(list)

	//res := answers.IsPalindromeList(list)
	//fmt.Println(res)

	//278. 第一个错误的版本
	//res := answers.FirstBadVersion(8)
	//fmt.Println(res)

	//350. 两个数组的交集 II
	//res := answers.Intersect([]int{1, 2, 3}, []int{3, 4, 5})
	//fmt.Printf("%v\n", res)

	//384. 打乱数组
	//obj := answers.Constructor([]int{1, 21, 3, 14, 5})
	//shuff := obj.Shuffle()
	//fmt.Printf("%v\n", shuff)
	//rest := obj.Reset()
	//fmt.Printf("%v\n", rest)

	//461. 汉明距离
	//res := answers.HammingDistance2(1, 4)
	//fmt.Println("distance is ", res)

	//468. 验证IP地址.验证IP是否合法
	//flags := answers.ValidIPAddress("192.0.0.1")
	//println(flags)

	//477. 汉明距离总和
	//res := answers.TotalHammingDistance([]int{4, 14, 2})
	//fmt.Println(res)

	//???523. 连续的子数组和
	//sum := answers.CheckSubarraySum([]int{23, 2, 4, 6, 7}, 6)
	//fmt.Println(sum)

	//???525. 连续数组
	//subarray := answers.FindMaxLength([]int{0, 1, 0})
	//fmt.Printf("%v\n", subarray)

	//???560. 和为K的子数组
	//count := answers.SubarraySum([]int{1, 1, 1}, 2)
	//fmt.Println(count)

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

	//387. 字符串中的第一个唯一字符
	//res := answers.FirstUniqChar("sasbadxf")
	//println(res)

	//剑指 Offer 53 - II. 0～n-1中缺失的数字
	//nums := []int{0, 1, 2, 3, 4, 5, 6, 7, 9}
	//res := answers.MissingNumber(nums)
	//println(res)

	//268. 丢失的数字
	//nums := []int{0, 1, 2, 3, 4, 5, 6, 7, 9}
	//res := answers.MissingNum(nums)
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

	//剑指 Offer 67. 把字符串转换成整数
	//intnum := answers.StrToInt("   -42")
	//fmt.Println(intnum)

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

	//143. 重排链表
	//res := answers.ReorderList()
	//dumpList(res)

	//144. 二叉树的前序遍历
	//94. 二叉树的中序遍历
	//145. 二叉树的后序遍历
	//103. 二叉树的锯齿形层序遍历
	//199. 二叉树的右视图
	root := &answers.TreeNode{
		Val: 5,
		Left: &answers.TreeNode{
			Val: 4,
			Left: &answers.TreeNode{
				Val:   1,
				Left:  nil,
				Right: nil,
			},
			Right: &answers.TreeNode{
				Val:   2,
				Left:  nil,
				Right: nil,
			},
		},
		Right: &answers.TreeNode{
			Val:   6,
			Left:  nil,
			Right: nil,
		},
	}
	//res := answers.PreorderTraversal(root)
	//res := answers.InorderTraversal(root)
	res := answers.PostorderTraversal(root)
	fmt.Printf("%v\n", res)

	//102. 二叉树的层序遍历
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
	//fmt.Printf("102 answer : %v\n", res)

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

	//287. 寻找重复数
	//var nums = []int{1, 2, 3, 4, 4}
	//dupNum := answers.FindDuplicate(nums)
	//fmt.Println(dupNum)

	//316. 去除重复字母
	//res := answers.RemoveDuplicateLetters("")
	//fmt.Println(res)

	//338. 比特位计数
	//countArr := answers.CountBits(7)
	//println(countArr)

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
	//				Val: 4,
	//				Next: &answers.ListNode{
	//					Val: 5,
	//					Next: &answers.ListNode{
	//						Val:  6,
	//						Next: nil,
	//					},
	//				},
	//			},
	//		},
	//	},
	//}
	//res := answers.GetKthFromEnd(head, 3)
	//dumpList(res)
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

//https://leetcode-cn.com/leetbook/read/top-interview-questions-easy/xn96us/
func isAnagram(s string, t string) bool {
	len1 := len(s)
	len2 := len(t)
	if len1 != len2 {
		return false
	}
	map1, map2 := make(map[byte]int), make(map[byte]int)
	for i := 0; i < len1; i++ {
		map1[s[i]]++
	}
	for i := 0; i < len1; i++ {
		map2[t[i]]++
	}

	for k1, v1 := range map1 {
		v2, ok := map2[k1]
		if !ok || v2 != v1 {
			return false
		}
	}

	return true
}

//https://leetcode-cn.com/leetbook/read/top-interview-questions-easy/xne8id/
func isPalindrome(s string) bool {
	var newStr = make([]byte, 0, len(s))

	for i := 0; i < len(s); i++ {
		if isLower(s[i]) {
			newStr = append(newStr, s[i])
		} else if isUpper(s[i]) {
			newStr = append(newStr, s[i]+32)
		} else if isNumber(s[i]) {
			newStr = append(newStr, s[i])
		}
	}
	//fmt.Printf("%s\n", newStr)

	size := len(newStr)
	start, end := 0, size-1
	for start < end {
		if newStr[start] != newStr[end] {
			return false
		}
		start++
		end--
	}

	return true
}

func isUpper(a byte) bool {
	return a >= 'A' && a <= 'Z'
}
func isLower(a byte) bool {
	return a >= 'a' && a <= 'z'
}
func isNumber(a byte) bool {
	return a >= '0' && a <= '9'
}

//https://leetcode-cn.com/leetbook/read/top-interview-questions-easy/xnoilh/
func myAtoi(s string) int {
	res, flag := 0, 1
	start, end := false, false
	for i := 0; i < len(s); i++ {
		switch s[i] {
		case '+':
			if start {
				end = true
			}
			start = true
		case '-':
			if start {
				end = true
			}
			start = true
		case ' ':
			if start {
				end = true
			}
			start = true
		case '1', '2', '3', '4', '5', '6', '7', '8', '9', '0':
			res += res*10 + int(s[i]-'0')
			start = true
		default:
			end = true
		}
		if end {
			break
		}
	}

	return res * flag
}

//题目 https://leetcode-cn.com/leetbook/read/top-interview-questions-easy/xnpvdm/
//解题 https://leetcode-cn.com/problems/count-and-say/solution/go-die-dai-by-ba-xiang/
func countAndSay(n int) string {
	str := "1"
	for i := 2; i <= n; i++ {
		temp := strings.Builder{}
		preByte := str[0]
		count := 1
		for j := 1; j < len(str); j++ {
			if preByte == str[j] {
				count++
			} else {
				temp.WriteString(strconv.Itoa(count))
				temp.WriteByte(preByte)
				count = 1
				preByte = str[j]
			}
		}
		//把最后的补上
		temp.WriteString(strconv.Itoa(count))
		temp.WriteByte(preByte)

		str = temp.String()
	}

	return str
}

//https://leetcode-cn.com/leetbook/read/top-interview-questions-easy/x2gy9m/
func removeDuplicateNodes(nums []int) int {
	size := len(nums)
	slow := 0
	for fast := 1; fast < size; fast++ {
		if nums[slow] != nums[fast] {
			slow++
			nums[slow] = nums[fast]
		}
	}

	if slow < size {
		return slow + 1
	}
	return size
}

//https://leetcode-cn.com/leetbook/read/top-interview-questions-easy/x2zsx1/
func maxProfit(prices []int) int {
	min_price := -1 << 31
	max_profit := 0
	for _, price := range prices {
		if price < min_price {
			min_price = price
		}
		if price-min_price > max_profit {
			max_profit = price - min_price
		}
	}
	return max_profit
}

//https://leetcode-cn.com/leetbook/read/top-interview-questions-easy/x2skh7/
//借助临时变量
func rotate(nums []int, k int) {
	size := len(nums)

	for k > size {
		k -= size
	}

	temp := make([]int, 0, size)
	start := size - k
	for i := start; i < size; i++ {
		temp = append(temp, nums[i])
	}
	for i := 0; i < start; i++ {
		temp = append(temp, nums[i])
	}

	copy(nums, temp)
}

//https://leetcode-cn.com/leetbook/read/top-interview-questions-easy/x2skh7/
//分三步 1.全部反转 2.反转前面部分 3.反转后面部分
func rotate2(nums []int, k int) {
	size := len(nums)

	for k > size {
		k %= size
	}

	//全部反转
	reverse(nums, 0, size-1)
	fmt.Printf("%v 11\n", nums)
	//反转前面部分
	reverse(nums, 0, k-1)
	fmt.Printf("%v 22\n", nums)
	//反转后面部分
	reverse(nums, k, size-1)
	fmt.Printf("%v 33\n", nums)
}
func reverse(nums []int, start, end int) {
	for start < end {
		nums[start], nums[end] = nums[end], nums[start]
		start++
		end--
	}
}

//https://leetcode-cn.com/leetbook/read/top-interview-questions-easy/x248f5/
func containsDuplicate(nums []int) bool {
	//先排序
	sort.Ints(nums)

	//在判断
	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i-1] {
			return true
		}
	}

	return false
}

//https://leetcode-cn.com/leetbook/read/top-interview-questions-easy/x21ib6/
func singleNumber(nums []int) int {
	reduce := 0
	for i := 0; i < len(nums); i++ {
		reduce ^= nums[i]
	}

	return reduce
}

//https://leetcode-cn.com/leetbook/read/top-interview-questions-easy/xngt85/
func fizzBuzz(n int) []string {
	list := make([]string, 0, n)
	for i := 1; i <= n; i++ {
		if i%3 == 0 {
			if i%5 == 0 {
				list = append(list, "FizzBuzz")
			} else {
				list = append(list, "Fizz")
			}
			continue
		}
		if i%5 == 0 {
			list = append(list, "Buzz")
			continue
		}
		list = append(list, strconv.Itoa(i))
	}

	return list
}

//https://leetcode-cn.com/leetbook/read/top-interview-questions-easy/xnzlu6/
func countPrimes(n int) int {
	isPrime := make([]bool, n)
	for i := range isPrime {
		isPrime[i] = true
	}

	count := 0
	for i := 2; i < n; i++ {
		if isPrime[i] {
			count++
			for j := 2; j*i < n; j++ {
				isPrime[j*i] = false
			}
		}
	}

	return count
}

func readMemStats() {

	var ms runtime.MemStats

	runtime.ReadMemStats(&ms)

	log.Printf(" ===> Alloc:%d(bytes) HeapIdle:%d(bytes) HeapReleased:%d(bytes)", ms.Alloc, ms.HeapIdle, ms.HeapReleased)
}

func test() {
	//slice 会动态扩容，用slice来做堆内存申请
	container := make([]int, 8)

	log.Println(" ===> loop begin.")
	for i := 0; i < 32*1000*1000; i++ {
		container = append(container, i)
		if i == 16*1000*1000 {
			readMemStats()
		}
	}

	log.Println(" ===> loop end.")
}
