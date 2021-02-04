package answers

//https://leetcode-cn.com/problems/latest-time-by-replacing-hidden-digits/
func MaximumTime(time string) string {
	var res [5]byte

	for i := 0; i < len(time); i++ {
		if time[i] == '?' {
			switch i {
			case 0:
				if time[i+1] == '?' || time[i+1] <= '3' {
					res[i] = '2'
				} else {
					res[i] = '1'
				}
			case 1:
				if time[i-1] == '?' || time[i-1] >= '2' {
					res[i] = '3'
				} else {
					res[i] = '9'
				}
			case 3:
				res[i] = '5'
			case 4:
				res[i] = '9'
			}
		} else {
			res[i] = time[i]
		}
	}

	return string(res[:])
}
