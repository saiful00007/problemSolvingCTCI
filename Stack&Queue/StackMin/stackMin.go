// Stack Min: How would you design a stack which, in addition to push and pop, has a function min
// which returns the minimum eiement? Push, pop and min should ail operate in 0(1 ) time.

package main

import "fmt"

type Stack struct {
	originalStack []int
	minStack      []int
}

func (s *Stack) Initialize() {
	s.originalStack = []int{}
	s.minStack = []int{}
}

func (s *Stack) Push(value int) {
	s.originalStack = append(s.originalStack, value)

	if len(s.minStack) == 0 || value <= s.minStack[len(s.minStack)-1] {
		s.minStack = append(s.minStack, value)
	}

}

func (s *Stack) Pop() int {
	if len(s.originalStack) == 0 {
		return 0
	}

	value := s.originalStack[len(s.originalStack)-1]
	s.originalStack = s.originalStack[:len(s.originalStack)-1]
	if value == s.minStack[len(s.minStack)-1] {
		s.minStack = s.minStack[:len(s.minStack)-1]
	}
	return value
}

func (s *Stack) Top() int {
	length := len(s.originalStack)

	if length == 0 {
		return -1
	}
	return s.originalStack[length-1]
}

func (s *Stack) StackMin() int {
	if len(s.minStack) > 0 {
		return s.minStack[len(s.minStack)-1]
	}
	return 0
}

//display method for stacks

func main() {
	newStack := Stack{}
	newStack.Push(-2)
	newStack.Push(0)
	newStack.Push(-3)

	fmt.Println(newStack.Top())
	fmt.Println(newStack.Pop())
	fmt.Println(newStack.StackMin())
}
