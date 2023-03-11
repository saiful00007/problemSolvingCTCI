package main

import "fmt"

type Node struct {
	Data int
	Next *Node
}

type LinkedList struct {
	Head *Node
	Tail *Node
}

func (l *LinkedList) PushBack(value int) {
	newNode := &Node{Data: value, Next: nil}
	if l.Head == nil {
		l.Head = newNode
		l.Tail = newNode
	} else {
		l.Tail.Next = newNode
		l.Tail = newNode
	}
}

func (l *LinkedList) Display() {
	if l.Head == nil {
		return
	}
	curr := l.Head
	for curr != nil {
		fmt.Printf("%v ->", curr.Data)
		curr = curr.Next
	}
	fmt.Println()
}

func (l *LinkedList) PartitionList(x int) *LinkedList {
	left := &Node{}
	right := &Node{}
	leftTail := left
	rightTail := right
	curr := l.Head
	for curr != nil {
		if curr.Data < x {
			leftTail.Next = curr
			leftTail = leftTail.Next
		} else {
			rightTail.Next = curr
			rightTail = rightTail.Next
		}
		curr = curr.Next
	}
	leftTail.Next = right.Next
	l.Head = left.Next
	l.Tail = rightTail
	rightTail.Next = nil

	return l
}

func main() {
	newList := LinkedList{Head: nil, Tail: nil}
	newList.PushBack(2)
	newList.PushBack(7)
	newList.PushBack(3)
	newList.PushBack(8)
	newList.PushBack(1)
	newList.PushBack(4)
	newList.Display()
	list := newList.PartitionList(4)
	list.Display()
}
