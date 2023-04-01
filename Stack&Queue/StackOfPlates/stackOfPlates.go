package main

import (
	"errors"
)

type SetOfStacks struct {
	capacity int
	stacks   [][]int
}

func NewSetOfStacks(capacity int) *SetOfStacks {
	return &SetOfStacks{
		capacity: capacity,
		stacks:   [][]int{[]int{}},
	}
}

func (s *SetOfStacks) Push(item int) {
	if len(s.stacks[len(s.stacks)-1]) == s.capacity {
		s.stacks = append(s.stacks, []int{})
	}
	s.stacks[len(s.stacks)-1] = append(s.stacks[len(s.stacks)-1], item)
}

func (s *SetOfStacks) Pop() (int, error) {
	if len(s.stacks) == 1 && len(s.stacks[0]) == 0 {
		return 0, errors.New("Set of Stacks is empty")
	}
	item := s.stacks[len(s.stacks)-1][len(s.stacks[len(s.stacks)-1])-1]
	s.stacks[len(s.stacks)-1] = s.stacks[len(s.stacks)-1][:len(s.stacks[len(s.stacks)-1])-1]
	if len(s.stacks[len(s.stacks)-1]) == 0 && len(s.stacks) > 1 {
		s.stacks = s.stacks[:len(s.stacks)-1]
	}
	return item, nil
}
