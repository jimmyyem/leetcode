package lib

import (
	"unicode/utf8"
)

func SubStrRune(s string, start, end int) string {
	//fmt.Println(utf8.RuneCountInString(s), len(s))
	if utf8.RuneCountInString(s) > len(s) {
		rs := []rune(s)
		return string(rs[start:end])
	}

	return s[start:end]
}
