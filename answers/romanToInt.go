package answers

func RomanToInt(s string) int {
	mp := map[string]int{
		"I":  1,
		"IV": 4,
		"V":  5,
		"IX": 9,
		"X":  10,
		"XL": 40,
		"L":  50,
		"XC": 90,
		"C":  100,
		"CD": 400,
		"D":  500,
		"CM": 900,
		"M":  1000,
	}

	var res int
	for i := 0; i < len(s); i++ {
		if i+1 < len(s) {
			if mp[string(s[i])] < mp[string(s[i+1])] {
				res += mp[string(s[i+1])] - mp[string(s[i])]
				i++
			} else {
				res += mp[string(s[i])]
			}
		} else {
			res += mp[string(s[i])]
		}
	}

	return res
}
