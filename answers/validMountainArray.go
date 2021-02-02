package answers

//https://leetcode-cn.com/problems/valid-mountain-array/
func ValidMountainArray(arr []int) bool {
	var middle int

	if len(arr) < 3 {
		return false
	}

	for i := 0; i+1 < len(arr); i++ {
		if arr[i] == arr[i+1] {
			return false
		}

		//递增，一直到中间节点
		if arr[i] < arr[i+1] && middle == 0 {

		} else {
			//转折节点不能是开头或者结尾
			if i == 0 || i == len(arr)-1 {
				return false
			}
			//记录转折节点
			if middle == 0 {
				middle = i
			}
			//转折节点后必须是递减
			if arr[i] <= arr[i+1] {
				return false
			}
		}
	}

	//转折节点必须在中间位置
	return middle > 0 && middle < len(arr)-1
}
