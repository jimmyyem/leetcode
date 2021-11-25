package answers

/**
https://leetcode-cn.com/problems/minimum-window-substring/
76. 最小覆盖子串
*/
func MinWindow(s string, t string) string {
	//初步判断
	if len(s) < len(t) {
		return ""
	}

	var freq [58]int
	var need [58]int
	unique := 0 //记录t字符串不同字符的个数
	//注：记录不同字符的目的是为了可以将多个字符当成一个字符看，
	//但这一个字符也有一个值就对于need某个字符的值，那是可以把多个字符看成一个字符
	//的最小值，当freq大于need对于的字符哈希值时，也同样可以看作一个字符
	//在缩小窗口分支中的 if freq[s[l]-'A'] < need[s[l]-'A']便有所体现
	res := "" //返回结果

	for i := 0; i < len(t); i++ {
		if need[t[i]-'A'] == 0 {
			unique++ //不同字符数量增加
		}
		need[t[i]-'A']++
	}

	curr := 0     //记录当前窗口里面包含t中不同字符的个数
	l, r := 0, -1 //窗口大小[l...r]

	for l < len(s) {
		if (r+1) < len(s) && unique != curr { //保证r在扩大窗口之后不会越界
			r++              //扩大窗口
			freq[s[r]-'A']++ //增大频率
			if freq[s[r]-'A'] == need[s[r]-'A'] {
				curr++ //若两个表的值想等，则curr加一
			}
		} else {
			freq[s[l]-'A'] --                    //减少频率
			if freq[s[l]-'A'] < need[s[l]-'A'] { //更新匹配条件curr
				curr--
			}
			l++ //缩小窗口
		}

		if curr == unique {
			//更新结果的两种情况：
			//1.结果仍然为空字符串
			//2.结果比当前成功匹配的窗口大小要大
			if res == "" || len(res) > (r-l+1) {
				res = s[l : r+1] //切片切取的方式来更新
			}
		}
	}

	return res
}
