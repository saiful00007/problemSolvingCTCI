// Remove Dups: Write code to remove duplicates from an unsorted linked list.
// FOLLOW UP
// How would you solve this problem if a temporary buffer is not allowed?

package main

import "fmt"

//node structure of linked list
type Node struct {
	Data int
	Next *Node
}

//linked list structure
type LinkedList struct {
	Head *Node
	Tail *Node
}

//pushback function
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

//method for display a linked list
func (l LinkedList) Display() {
	curr := l.Head
	for curr != nil {
		fmt.Printf("%v ->", curr.Data)
		curr = curr.Next
	}
	fmt.Println()
}

//remove dups by recieving linked list
func (l *LinkedList) RemoveDups() {
	if l.Head == nil {
		return
	}
	curr := l.Head
	prev := l.Head
	unique := make(map[int]bool)

	for curr != nil {
		if _, ok := unique[curr.Data]; ok {
			prev.Next = curr.Next
		} else {
			unique[curr.Data] = true
			prev = curr
		}
		curr = curr.Next
	}
}

func main() {
	//Create a new list
	newList := LinkedList{Head: nil, Tail: nil}
	newList.PushBack(10)
	newList.PushBack(10)
	newList.PushBack(30)
	newList.PushBack(30)
	newList.PushBack(40)
	newList.PushBack(50)
	newList.Display()
	newList.RemoveDups()
	newList.Display()
}
