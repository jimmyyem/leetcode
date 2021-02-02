package answers

import "strings"

func LongestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	var res string
	first := strs[0]

	for i := 1; i < len(first); i++ {
		prefix := first[:i]

		for _, str := range strs {
			if strings.Index(str, prefix) == -1 {
				return res
			}
		}
		res = prefix
	}

	return res
}
