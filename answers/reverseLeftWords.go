package answers

func ReverseLeftWords(s string, n int) string {
	if n < 1 {
		return ""
	}

	var left, right string
	left = s[0:n]
	right = s[n:]

	return right + left
}
