// Sort Stack: Write a program to sort a stack such that the smallest items are on the top. You can use 
// an additional temporary stack, but you may not copy the elements into any other data structure 
// (such as an array). The stack supports the following operations: push, pop, peek, and is Empty.

package main

import "fmt"

type Stack struct {
	arrayStack []int
}

func (s *Stack) Len() int {
	return len(s.arrayStack)
}

func (s *Stack) IsEmpty() bool {
	return s.Len() == 0
}

func (s *Stack) Push(value int) {
	s.arrayStack = append(s.arrayStack, value)
}

func (s *Stack) Pop() int {
	if s.IsEmpty() {
		return 0
	}
	lastIndex := s.Len() - 1
	value := s.arrayStack[lastIndex]
	s.arrayStack = s.arrayStack[:lastIndex]
	return value
}

func (s *Stack) Sort() {
	if s.IsEmpty() {
		return
	}
	tempStack := Stack{}

	for !s.IsEmpty() {
		value := s.Pop()

		for !tempStack.IsEmpty() && tempStack.arrayStack[tempStack.Len()-1] > value {
			poppedValue := tempStack.Pop()
			s.Push(poppedValue)
		}
		tempStack.Push(value)
	}
	for !tempStack.IsEmpty() {
		value := tempStack.Pop()
		s.Push(value)
	}
}

func main() {
	newStack := Stack{}
	newStack.Push(5)
	newStack.Push(3)
	newStack.Push(6)
	newStack.Push(2)
	newStack.Push(1)
	newStack.Sort()
	for i := newStack.Len(); i > 0; i-- {
		fmt.Println(newStack.Pop())
	}
}
