package main

import (
	"fmt"
	"strings"
)

// 修复代码问题
func Solution(N int) {
	var enable_print int
	enable_print = N % 10
	for N > 0 {
		if enable_print == 0 && N%10 != 0 {
			enable_print = 1
			fmt.Print(N % 10)
		} else if enable_print > 0 {
			fmt.Print(N % 10)
		}
		N = N / 10
	}
}

func Solution2(A []int) int {
	// write your code in Go 1.4
	size := len(A)
	preSum := make([]int, size+1)
	for i := 1; i < size; i++ {
		// preSum[0]是固定的
		preSum[i] = preSum[i-1] + A[i-1]
	}
	//fmt.Printf("%v\n", A)
	//fmt.Printf("%v\n", preSum)

	res := 0
	for i := 0; i < size; i++ {
		for j := i; j < size; j++ {
			// preSum[3]-preSum[2]==0 说明肯定是中间有0元素或者和为0的子数组了
			if preSum[j+1]-preSum[i] == 0 {
				res++
			}
		}
	}

	return res
}

func Solution3(A []int, B []int, N int) int {
	connections := make([][]bool, N)
	for i := 0; i < N; i++ {
		connections[i] = make([]bool, N)
	}
	cache := make([]int, N)

	for i := 0; i < len(A); i++ {
		cache[A[i]]++
		cache[B[i]]++
		connections[A[i]][B[i]] = true
		connections[B[i]][A[i]] = true
	}
	res := 0
	for i := 0; i < N; i++ {
		for j := i + 1; j < N; j++ {
			if connections[i][j] {
				if res < cache[i]+cache[j]-1 {
					res = cache[i] + cache[j] - 1
				}
			} else {
				if res < cache[i]+cache[j] {
					res = cache[i] + cache[j]
				}
			}
		}
	}
	return res
}

type City struct {
	Name string
	Lin  []string
	Len  int
}

func loop(areas []string) {
	// 遍历所有的洲，做一个数组记录所有洲与相邻的洲的对应关系
	continents := map[string]City{} //洲名=>所有相邻的洲

	// 找到相邻的洲
	findArea := func(area string) City {
		city := City{}

	}

	for _, area := range areas {
		if _, ok := continents[area]; !ok {
			continents[area] = City{}
		}
		// 找到相邻的洲
		lin := findArea(area)
		continents[area] = lin
	}

	// 按相邻洲的个数排序，倒叙

	// 遍历这个数组，

}

func maxValue(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func minValue(a, b int) int {
	if a > b {
		return b
	}
	return a
}

// https://app.codility.com/test/589F9W-5J4/
// Tesla interview 2021-12-10 11:55:04
func main() {
	//res := Solution3([]int{1, 2, 4, 5}, []int{2, 3, 5, 6}, 6)
	res := Solution2([]int{2, -2, 3, 0, 4, -7})
	println(res)
	//Solution(10000)
	//println(" => 10000")
	//Solution(10005)
	//println(" => 10005")
	//Solution(100100)
	//println(" => 100100")
	//Solution(12345)
	//println(" => 12345")
	//Solution(100010020)
	//println(" => 10001002")
}

func commonPrefix(first, second string) int {
	if len(first) == 0 && len(second) != 0 {
		return len(second)
	}

	for i := 1; i < len(first); i++ {
		if !strings.HasPrefix(second, first[:i]) {
			return i - 1
		}
	}

	return 0
}
