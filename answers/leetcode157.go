package answers

//https://leetcode-cn.com/problems/read-n-characters-given-read4/s
//157. 用 Read4 读取 N 个字符

/**
 * The read4 API is already defined for you.
 *
 *     read4 := func(buf4 []byte) int
 *
 * // Below is an example of how the read4 API can be called.
 * file := File("abcdefghijk") // File is "abcdefghijk", initially file pointer (fp) points to 'a'
 * buf4 := make([]byte, 4) // Create buffer with enough space to store characters
 * read4(buf4) // read4 returns 4. Now buf = ['a','b','c','d'], fp points to 'e'
 * read4(buf4) // read4 returns 4. Now buf = ['e','f','g','h'], fp points to 'i'
 * read4(buf4) // read4 returns 3. Now buf = ['i','j','k',...], fp points to end of file
 */

var solution = func(read4 func([]byte) int) func([]byte, int) int {
	// implement read below.
	return func(buf []byte, n int) int {
		sum := 0 //读取的内容总长度
		res := make([]byte, 4)

		for sum < n {
			num := read4(res)

			//如果读到了内容
			if num > 0 {
				left := n - sum
				if left >= 4 || num <= left { //剩余2个 取到2个，剩余3个 取到2个
					copy(buf[sum:], res[:])
					sum += num
				} else if left < num {
					//剩余2个，但是取到3个
					copy(buf[sum:], res[:left])
					sum += left
				}
			}

			//如果读取够了，或者没内容可读取了
			if sum >= n || num == 0 {
				break
			}
		}

		return sum
	}
}
