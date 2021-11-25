package lib

import "math/rand"

/**
 * @see https://colobu.com/2018/09/02/generate-random-string-in-Go/
 */

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const (
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits
)

// 生成随机字符串-根据随机数、掩码位运算
func GetRandString(n int) string {
	b := make([]byte, n)
	size := len(letterBytes)
	for i := 0; i < n; {
		if num := int(rand.Int63() & letterIdxMask); num < size {
			b[i] = letterBytes[num]
			i++
		}
	}

	return string(b)
}

// 生成随机字符串-根据随机数、掩码位运算- 升级版
func GetRandStringMask(n int) string {
	b := make([]byte, n)
	for i, cache, remain := n-1, rand.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = rand.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}

	return string(b)
}
