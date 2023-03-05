//  Delete Middle Node: Implement an algorithm to delete a node in the middle (i.e., any node but
// 	the first and last node, not necessarily the exact middle) of a singly linked list,
// 	given only access to that node.
// 	EXAMPLE
// 	Input: the node c from the linked list a - >b- >c - >d - >e- >f
// 	Result: nothing is returned, but the new linked list looks like a->b->d->e-> f

package main

import (
	"fmt"
)

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

func (l LinkedList) Display() {
	curr := l.Head
	for curr != nil {
		fmt.Printf("%v ->", curr.Data)
		curr = curr.Next
	}
	fmt.Println()
}

func (l *LinkedList) deleteMiddleNode() bool {
	if l.Head == nil || l.Head.Next == nil {
		//linked list is empty or has only one node only
		return false
	}
	//tow pointers for travers the linked list
	slow := l.Head
	fast := l.Head
	//this pointer is for keep track of the previous node of slow
	prev := l.Head
	//traverse the liked list and find the middle node
	for fast != nil && fast.Next != nil {
		prev = slow
		slow = slow.Next
		fast = fast.Next.Next
	}

	if prev == nil {
		// The length of the list is even
		l.Head = slow.Next
	} else {
		prev.Next = slow.Next
	}
	return true
}

func main() {
	newList := LinkedList{}
	newList.PushBack(10)
	newList.PushBack(20)
	newList.PushBack(30)
	newList.PushBack(40)
	newList.PushBack(50)
	fmt.Println("Original Linked List")
	newList.Display()
	if newList.deleteMiddleNode() {
		fmt.Println("List after middle node deleted")
		newList.Display()
	} else {
		fmt.Println("List is Empty")
	}
}
