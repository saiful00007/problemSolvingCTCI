//  Describe how you could use a single array to implement three stacks.

package main

import (
	"fmt"
)

type ThreeStacks struct {
	arrayStack []int
	top1       int
	top2       int
	top3       int
	capacity   int
}

// create a new instance of the threestacks struct with the specific capacity
func NewThreeStacks(capacity int) *ThreeStacks {
	return &ThreeStacks{
		arrayStack: make([]int, 3*capacity),
		top1:       -1,
		top2:       capacity - 1,
		top3:       capacity*2 - 1,
		capacity:   capacity,
	}
}

// push method for stacks individually works
func (ts *ThreeStacks) Push(stkNumber int, value int) {
	switch stkNumber {
	case 1:
		if ts.top1+1 >= ts.capacity {
			fmt.Println("stack1 overflow")
			return
		}
		ts.top1++
		ts.arrayStack[ts.top1] = value
	case 2:
		if ts.top2+1 >= ts.capacity*2 {
			fmt.Println("stack2 overflow")
			return
		}
		ts.top2++
		ts.arrayStack[ts.top2] = value
	case 3:
		if ts.top3+1 >= ts.capacity*3 {
			fmt.Println("Stack3 Overflow")
			return
		}
		ts.top3++
		ts.arrayStack[ts.top3] = value
	default:
		fmt.Println("Invalid stack number")
	}
}

// pop method for three stacks
func (ts *ThreeStacks) Pop(stkNum int) int {
	switch stkNum {
	case 1:
		//check if the stack is empty
		if ts.top1 == -1 {
			fmt.Println("stack1 underflow")
			return -1
		}
		value := ts.arrayStack[ts.top1]
		ts.top1--
		return value
	case 2:
		//check if the stack is empty
		if ts.top2 == ts.capacity-1 {
			fmt.Println("stack2 underflow")
			return -1
		}
		value := ts.arrayStack[ts.top2]
		ts.top2--
		return value
	case 3:
		//check if the stack is empty
		if ts.top3 == ts.capacity*2 {
			fmt.Println("stack3 underflow")
			return -1
		}
		value := ts.arrayStack[ts.top3]
		ts.top3--
		return value
	default:
		fmt.Println("invalid stack number")
		return -1
	}
}

// create top method for top element of the specific stack
func (ts *ThreeStacks) Top(stkNum int) int {
	switch stkNum {
	case 1:
		if ts.top1 == -1 {
			fmt.Println("stack1 is empty")
			return -1
		}
		return ts.arrayStack[ts.top1]
	case 2:
		if ts.top2 == ts.capacity-1 {
			fmt.Println("stack2 is empty")
			return -1
		}
		return ts.arrayStack[ts.top2]
	case 3:
		if ts.top3 == ts.capacity*2 {
			fmt.Println("stack3 is empty")
			return -1
		}
		return ts.arrayStack[ts.top3]
	default:
		fmt.Println("Invalid stack number")
		return -1
	}
}

// create display method for stack display
func (ts *ThreeStacks) DisplayStack() {
	fmt.Println("Stack 1:")
	for i := ts.top1; i >= 0; i-- {
		fmt.Println(ts.arrayStack[i])
	}
	fmt.Println("Stack 2:")
	for i := ts.top2; i >= ts.capacity; i-- {
		fmt.Println(ts.arrayStack[i])
	}
	fmt.Println("Stack 3:")
	for i := ts.top3; i >= ts.capacity*2; i-- {
		fmt.Println(ts.arrayStack[i])
	}
}

func main() {
	newThreeStacks := NewThreeStacks(5)

	newThreeStacks.Push(1, 10)
	newThreeStacks.Push(1, 20)
	newThreeStacks.Push(1, 30)
	newThreeStacks.Push(1, 40)
	newThreeStacks.Push(2, 10)
	newThreeStacks.Push(1, 40)
	newThreeStacks.Push(1, 40)
	newThreeStacks.Push(1, 40)
	newThreeStacks.Push(1, 40)
	newThreeStacks.Push(2, 20)
	newThreeStacks.Push(2, 30)
	newThreeStacks.Push(2, 40)
	newThreeStacks.Push(3, 10)
	newThreeStacks.Push(3, 20)
	newThreeStacks.Push(1, 40)
	newThreeStacks.Push(1, 40)
	newThreeStacks.Push(1, 40)
	newThreeStacks.Push(3, 30)
	newThreeStacks.Push(3, 40)

	newThreeStacks.DisplayStack()

	newThreeStacks.Pop(1)
	newThreeStacks.Pop(2)
	newThreeStacks.Pop(3)

	fmt.Println(newThreeStacks.arrayStack)
	newThreeStacks.DisplayStack()
}
