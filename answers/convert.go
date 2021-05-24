package answers

import (
	"strings"
	//"fmt"
)

//https://leetcode-cn.com/problems/zigzag-conversion/
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
