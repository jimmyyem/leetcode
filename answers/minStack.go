package answers

import "math"

//https://leetcode-cn.com/problems/min-stack/
//155. 最小栈

type MinStack struct {
	stack    []int
	minStack []int
}

/** initialize your data structure here. */
func Constructor2() MinStack {
	return MinStack{
		stack:    []int{},
		minStack: []int{math.MaxInt64},
	}
}

//push
func (this *MinStack) Push(val int) {
	this.stack = append(this.stack, val)

	currMinVal := this.minStack[len(this.minStack)-1]
	this.minStack = append(this.minStack, min(val, currMinVal))
}

//pop
func (this *MinStack) Pop() {
	this.stack = this.stack[:len(this.stack)-1]
	this.minStack = this.minStack[:len(this.minStack)-1]
}

//top
func (this *MinStack) Top() int {
	return this.stack[len(this.stack)-1]
}

//getMin
func (this *MinStack) GetMin() int {
	return this.minStack[len(this.minStack)-1]
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
