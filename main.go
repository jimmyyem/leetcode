package main

import (
	"fmt"
	"leetcode/answers"
)

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

	fmt.Println("重新排列日志文件")
	arr := []string{"a1 9 2 3 1","g1 act car","zo4 4 7","ab1 off key dog","a8 act zoo"}
	res := answers.ReorderLogFiles(arr)
	fmt.Println(res)
}
