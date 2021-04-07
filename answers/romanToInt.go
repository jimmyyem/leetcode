package answers

//https://leetcode-cn.com/problems/roman-to-integer/
//13. 罗马数字转整数
/**
仔细阅题 会发现如下规律：
XXVII等于1+1+5+10+10 = 27 、IX等于10-1=9、XCI等于1+100-10。

思路：从右往左遍历，当前的字符大于等于上一个字符，则累加； 否则累减

作者：dempsey-2
链接：https://leetcode-cn.com/problems/roman-to-integer/solution/luo-ma-zi-fu-zhuan-zheng-shu-by-dempsey-2/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
 */
func RomanToInt(s string) int {
	var _map = map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	// XXVII等于1+1+5+10+10 = 27 、IX等于10-1=9、XCI等于1+100-10。

	// 当右边的字符比左边的大，说明是特殊组合
	pre, r := 0,0
	for i:=len(s)-1; i>=0;i-- {
		cur := _map[s[i]]
		if cur >= pre {
			r += cur
		}else {
			r -= cur
		}
		pre = cur
	}
	return r
}
