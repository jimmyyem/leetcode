package answers

import "fmt"

//https://leetcode-cn.com/problems/implement-strstr/
//28. 实现 strStr()
//Brute Force  https://github.com/chefyuan/algorithm-base/blob/main/animation-simulation/%E6%95%B0%E6%8D%AE%E7%BB%93%E6%9E%84%E5%92%8C%E7%AE%97%E6%B3%95/BF%E7%AE%97%E6%B3%95.md
func StrStrBF(s string, substr string) int {
	n, m := len(s), len(substr)

	if m == 0 {
		return 0
	}
	if n < m || n == 0 {
		return -1
	}

	for i := 0; i <= n-m; i++ {
		a, b := i, 0
		for b < m && s[a] == substr[b] {
			a++
			b++
		}
		if b == m {
			return i
		}
	}
	return -1
}

func StrStr(s, substr string) int {
	n, m := len(s), len(substr)
	for i := 0; i <= n-m; i++ {
		a, b := i, 0
		for b < m && s[a] == substr[b] {
			a++
			b++
		}
		if b == m {
			return i
		}
	}
	return -1
}

//寻找子串在s中首次出现的位置
//BM 算法也是由两位发明者 Boyer 和 Moore 的名字命名的，其核心思想是在主串和模式串进行比较时，如果出现了不匹配的情况，能够尽可能多的获取一些信息，借此跳过那些肯定不会匹配的情况，以此来提升字符串匹配的效率，大多数文本编辑器中的字符查找功能，一般都会使用 BM 算法来实现。
func StrStrBM(s, substr string) int {
	var len1, len2 = len(s), len(substr)

	if len2 == 0 {
		return 0
	}
	if len1 == 0 || len1 < len2 {
		return -1
	}
	if s == substr {
		return 0
	}

	return -1
}

func StrStrKMP(haystack string, needle string) int {
	if len(needle) == 0 { //若模式串为空串
		return 0
	}
	if len(haystack) == 0 || len(haystack) < len(needle) {
		return -1
	}

	next := getNextRes(needle)
	//fmt.Printf("%v\n", next)
	max_match := 0                                   //当前最大匹配长度
	for index := 0; index < len(haystack); index++ { //开始匹配
		for max_match > 0 && haystack[index] != needle[max_match] {
			max_match = next[max_match-1] //回溯
		}
		if haystack[index] == needle[max_match] {
			max_match++
		}

		if max_match == len(needle) { //如果max_match长度到达模式串的长度则说明匹配完成
			return index - (max_match - 1)
		}
	}

	return -1
}
func getNextRes(s string) []int {
	next := make([]int, len(s))
	max_match := 0 //当前最大匹配长度
	for index := 1; index < len(s); index++ {

		for max_match > 0 && s[index] != s[max_match] {
			max_match = next[max_match-1] //回溯
		}
		if s[index] == s[max_match] {
			max_match++
			next[index] = max_match
		}

	}
	return next
}

func KMP(haystack, needle string) int {
	fixArray := make([]int, len(needle))
	for i := 2; i < len(needle); i++ {
		fixArray[i] = QuerySameFix(haystack[:i])
		//fmt.Printf("%s => %d\n", haystack[:i], fixArray[i])
	}
	fmt.Printf("%v\n", fixArray)

	return -1
}

//获得一个字符串的相同前后缀
func QuerySameFix(s string) int {
	mid, length := len(s)/2, len(s)
	//fmt.Println(mid, length)
	for i := mid; i > 0; i-- {
		if s[:i] != s[length-i:] {
			return i
		}

	}

	return 0
}
