package answers

//https://leetcode-cn.com/problems/ti-huan-kong-ge-lcof/
//剑指 Offer 05. 替换空格
func ReplaceSpace(s string) string {
	var res = make([]byte, len(s)*3)
	var i, idx = 0, 0

	for ; i < len(s); i++ {
		if s[i] == ' ' {
			res[idx] = '%'
			idx++
			res[idx] = '2'
			idx++
			res[idx] = '0'
			idx++
		} else {
			res[idx] = s[i]
			idx += 1
		}
	}

	return string(res[:idx])

	//var res strings.Builder
	//for i := range s {
	//	if s[i] == ' ' {
	//		res.WriteString("%20")
	//	} else {
	//		res.WriteByte(s[i])
	//	}
	//}
	//
	//return res.String()
}
