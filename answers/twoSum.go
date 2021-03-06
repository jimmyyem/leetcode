package answers

//func main() {
//	nums := []int{3, 2, 4}
//	target := 6
//	res := twoSum(nums, target)
//	fmt.Println(res)
//
//}

// [3,2,4],  9  返回[1,2]
//https://leetcode-cn.com/problems/two-sum/
//1.两数之和
func TwoSum(nums []int, target int) []int {
	mp := make(map[int]int, len(nums))
	for idx, value := range nums {
		diffValue := target - value
		diffIndex, found := mp[diffValue]
		if found {
			return []int{idx, diffIndex}
		} else {
			mp[value] = idx
		}
	}

	return nil
}
