package main

import "fmt"

type HeapNode struct {
	Val         int
	Left, Right *HeapNode
}

//堆的数组存储
func main() {
	head := &HeapNode{
		Val: 0,
		Left: &HeapNode{
			Val: 1,
			Left: &HeapNode{
				Val:   3,
				Left:  nil,
				Right: nil,
			},
			Right: &HeapNode{
				Val:   4,
				Left:  nil,
				Right: nil,
			},
		},
		Right: &HeapNode{
			Val: 2,
			Left: &HeapNode{
				Val:   5,
				Left:  nil,
				Right: nil,
			},
			Right: nil,
		},
	}
	dumpHeap(head)
	nodeArray := toArray(head)
	fmt.Printf("%v\n", nodeArray)

	nums := []int{1, 5, 9, 3, 4, 2, 6, 8, 22}
	heapsort(nums)
}

//堆转数组存储
func toArray(head *HeapNode) []int {
	res := make([]int, 1)
	res[0] = -1
	nodes := []*HeapNode{head}
	temp := make([]*HeapNode, 0)
	for {
		for i := 0; i < len(nodes); i++ {
			fmt.Printf("%d ", nodes[i].Val)
			res = append(res, nodes[i].Val)
			if nodes[i].Left != nil {
				temp = append(temp, nodes[i].Left)
			}
			if nodes[i].Right != nil {
				temp = append(temp, nodes[i].Right)
			}
		}
		//遍历结束
		if len(temp) == 0 {
			break
		}
		nodes = temp
		temp = temp[:0]
	}

	return res
}

//打印堆内容
func dumpHeap(head *HeapNode) {
	nodes := []*HeapNode{head}
	temp := make([]*HeapNode, 0)
	for {
		for i := 0; i < len(nodes); i++ {
			fmt.Printf("%d ", nodes[i].Val)
			if nodes[i].Left != nil {
				temp = append(temp, nodes[i].Left)
			}
			if nodes[i].Right != nil {
				temp = append(temp, nodes[i].Right)
			}
		}
		//遍历结束
		if len(temp) == 0 {
			break
		}
		nodes = temp
		temp = temp[:0]
	}
	fmt.Println()
	return
}

//升序
func heapsort(nums []int) {
	fmt.Printf("初始状态 %v\n", nums)
	n := len(nums) - 1 //nums[0]不用
	for k := n / 2; k >= 1; k-- {
		sink(nums, k, n)
	}
	fmt.Printf("建堆后 %v\n", nums)

	for n > 1 {
		swap(nums, 1, n)
		n--
		sink(nums, 1, n)
	}
	fmt.Printf("排序后 %v", nums)
}

func sink(nums []int, k, n int) {
	for {
		i := k * 2
		if i > n {
			break
		}
		if i < n && nums[i+1] > nums[i] {
			i++
		}
		if nums[k] >= nums[i] {
			break
		} else {
			swap(nums, k, i)
		}

		k = i
	}
}

func swap(nums []int, i, j int) {
	nums[i], nums[j] = nums[j], nums[i]
}
