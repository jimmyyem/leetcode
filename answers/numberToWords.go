package answers

import "strings"

//https://leetcode-cn.com/problems/integer-to-english-words/
//273.整数转换英文表示
func NumberToWords(num int) string {
	if num == 0 {
		return "Zero"
	}

	pair := map[int]string{
		0:          "Zero",
		1:          "One",
		2:          "Two",
		3:          "Three",
		4:          "Four",
		5:          "Five",
		6:          "Six",
		7:          "Seven",
		8:          "Eight",
		9:          "Nine",
		10:         "Ten",
		11:         "Eleven",
		12:         "Twelve",
		13:         "Thirteen",
		14:         "Fourteen",
		15:         "Fifteen",
		16:         "Sixteen",
		17:         "Seventeen",
		18:         "Eighteen",
		19:         "Nineteen",
		20:         "Twenty",
		30:         "Thirty",
		40:         "Forty",
		50:         "Fifty",
		60:         "Sixty",
		70:         "Seventy",
		80:         "Eighty",
		90:         "Ninety",
		100:        "Hundred",
		1000:       "Thousand",
		1000000:    "Million",
		1000000000: "Billion",
	}

	transfer100 := func(num int) string {
		res := make([]string, 0)
		for num > 0 {
			if num <= 20 {
				res = append(res, pair[num])
				break
			}

			if num < 100 {
				whole := num / 10 * 10
				num %= 10
				res = append(res, pair[whole])
			} else {
				whole := num / 100
				res = append(res, pair[whole])
				res = append(res, pair[100])
				num %= 100
			}
		}

		return strings.Join(res, " ")
	}

	var res = make([]string, 0)

	for num > 0 {
		if num >= 1000000000 {
			whole := num / 1000000000
			res = append(res, pair[whole])
			res = append(res, pair[1000000000])
			num %= 1000000000
		} else if num >= 1000000 {
			whole := num / 1000000
			res = append(res, transfer100(whole))
			res = append(res, pair[1000000])
			num %= 1000000
		} else if num >= 1000 {
			whole := num / 1000
			res = append(res, transfer100(whole))
			res = append(res, pair[1000])
			num %= 1000
		} else if num >= 100 {
			whole := num / 100
			res = append(res, transfer100(whole))
			res = append(res, pair[100])
			num %= 100
		} else {
			res = append(res, transfer100(num))
			break
		}
	}

	return strings.Trim(strings.Join(res, " "), " ")
}
