package answers

var (
	cache = make(map[int]int)
)

func Fabi(n int, useCache bool) int {
	if useCache {
		if val, ok := cache[n]; ok {
			return val
		}
	}

	switch n {
	case 0:
		return 0
	case 1, 2:
		return 1
	default:
		res := Fabi(n-1, useCache) + Fabi(n-2, useCache)
		if useCache {
			cache[n] = res
		}
		return res
	}
}
