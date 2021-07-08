package answers

import (
	"strconv"
	"strings"
)

//https://leetcode-cn.com/problems/validate-ip-address/
//468. 验证IP地址
func ValidIPAddress(IP string) string {
	var res string

	if strings.Contains(IP, ":") {
		res = validIPv6(IP)
	} else {
		res = validIPv4(IP)
	}

	return res
}

//ipv6
func validIPv6(IP string) string {
	numSets := strings.Split(IP, ":")

	// 排除格式非法
	if len(numSets) != 8 {
		return "Neither"
	}

	for _, num := range numSets {
		// 排除 ::
		if len(num) == 0 || len(num) > 4 {
			return "Neither"
		}

		// 根据ascii码排除非16进制数
		for i := 0; i < len(num); i++ {
			if (num[i] > 'F' && num[i] < 'a') || num[i] > 'f' || num[i] < '0' {
				return "Neither"
			}
		}
	}

	return "IPv6"
}

//ipv4
func validIPv4(IP string) string {
	numSets := strings.Split(IP, ".")

	//排除格式非法
	if len(numSets) != 4 {
		return "Neither"
	}

	for _, num := range numSets {
		//排除空字符串
		if len(num) == 0 {
			return "Neither"
		}

		// 处理 01
		if len(num) > 1 && num[0] == '0' {
			return "Neither"
		}
		number, err := strconv.Atoi(num)

		//非法数字，处理 266
		if err != nil || number > 255 {
			return "Neither"
		}
	}

	return "IPv4"
}

