// Return Kth to Last: Implement an algorithm to find the kth to last element of a singly linked list.
// iterate the list with 2 variables..one is used for finding the kth position of the list
// then
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

func (l *LinkedList) KthToLast(k int) *Node {
	if l.Head == nil || k < 0 {
		return nil
	}
	curr := l.Head
	for i := 0; i < k; i++ {
		if curr == nil {
			return nil
		}
		curr = curr.Next
	}
	nthList := l.Head
	for curr != nil {
		nthList = nthList.Next
		curr = curr.Next
	}
	return nthList
}

func (l LinkedList) Display() {
	for l.Head.Data == 0 {
		return
	}
	curr := l.Head
	for curr != nil {
		fmt.Printf("%v ->", curr.Data)
		curr = curr.Next
	}
	fmt.Println()
}

func main() {
	newList := LinkedList{Head: nil, Tail: nil}
	newList.PushBack(10)
	newList.PushBack(20)
	newList.PushBack(30)
	newList.PushBack(40)
	newList.PushBack(50)
	newList.Display()
	k := 2
	kth := newList.KthToLast(k)
	fmt.Printf("the %dnd to last element is %d\n ->", k, kth.Data)
}
