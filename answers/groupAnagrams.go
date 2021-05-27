package answers

import "sort"

//https://leetcode-cn.com/problems/group-anagrams/
//49. 字母异位词分组
func GroupAnagrams(strs []string) [][]string {
	res := [][]string{}	//结果
	m := make(map[string][]string, len(strs))	//保存临时结果的map

	for _, v := range strs {
		//string转成[]byte，用于排序
		s := []byte(v)
		sort.Slice(s, func(i, j int) bool { return s[i] < s[j] }) // 排序
		//排序后的字符串作为key
		key := string(s) // 异位词排序后一定是相同的
		// map的键是排序后的字符串,值是字符串切片,用于存储分组
		m[key] = append(m[key], v)
	}
	for _, v := range m {
		res = append(res, v)
	}
	return res
}
