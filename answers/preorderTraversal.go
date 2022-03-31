package answers

import "errors"

// https://leetcode-cn.com/problems/binary-tree-preorder-traversal/
// 144. 二叉树的前序遍历（当前节点，左节点，右节点）
// 递归法 实现前序遍历
func PreorderTraversal(root *TreeNode) (res []int) {
	if root == nil {
		return []int{}
	}

	var preorder func(*TreeNode)
	preorder = func(root *TreeNode) {
		if root == nil {
			return
		}
		res = append(res, root.Val) //当前节点
		preorder(root.Left)         //左节点
		preorder(root.Right)        //右节点
	}

	preorder(root)
	return
}

// 迭代法 实现前序遍历
func preorderTraversal(root *TreeNode) (res []int) {
	if root == nil {
		return []int{}
	}

	stack := NewStackNode()
	stack.Push(root)

	for !stack.IsEmpty() {
		pop, _ := stack.Pop()
		node := pop.(*TreeNode)
		res = append(res, node.Val)

		if node.Right != nil {
			stack.Push(node.Right)
		}
		if node.Left != nil {
			stack.Push(node.Left)
		}
	}

	return
}

type StackNode []interface{}

func (stack StackNode) Len() int {
	return len(stack)
}

func (stack StackNode) IsEmpty() bool {
	return len(stack) == 0
}

func (stack StackNode) Cap() int {
	return cap(stack)
}

func (stack *StackNode) Push(value interface{}) {
	*stack = append(*stack, value)
}

func (stack StackNode) Top() (interface{}, error) {
	if len(stack) == 0 {
		return nil, errors.New("Out of index, len is 0")
	}
	return stack[len(stack)-1], nil
}

func (stack *StackNode) Pop() (interface{}, error) {
	theStack := *stack
	if len(theStack) == 0 {
		return nil, errors.New("Out of index, len is 0")
	}
	value := theStack[len(theStack)-1]
	*stack = theStack[:len(theStack)-1]
	return value, nil
}

func NewStackNode() *StackNode {
	return &StackNode{}
}
