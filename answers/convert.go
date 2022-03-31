package answers

import (
	"strings"
	//"fmt"
)

//https://leetcode-cn.com/problems/zigzag-conversion/
//https://leetcode-cn.com/problems/zigzag-conversion/solution/ren-zhe-suan-fa-chao-qing-xi-yi-dong-bai-40yv/
//6. Z 字形变换
func Convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}

	var res = make([]string, numRows)
	var period = numRows*2 - 2
	for i := 0; i < len(s); i++ {
		var mod = i % period
		//fmt.Println(mod, period)
		if mod < numRows {
			res[mod] += string(s[i])
		} else {
			res[period-mod] += string(s[i])
		}
	}
	//fmt.Printf("%v\n", res)
	return strings.Join(res, "")
}
