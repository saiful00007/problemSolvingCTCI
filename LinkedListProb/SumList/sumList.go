// Sum Lists: You have two numbers represented by a linked list, where each node contains a single
// digit. The digits are stored in reverse order, such that the Vs digit is at the head of the list. Write a
// function that adds the two numbers and returns the sum as a linked list.
// EXAMPLE
// Input: (7- > 1 -> 6) + (5 -> 9 -> 2).That is,617 + 295.
// Output: 2 -> 1 -> 9. That is, 912.
// FOLLOW UP
// Suppose the digits are stored in forward order. Repeat the above problem.
// EXAMPLE
// Input: (6 -> 1 -> 7) + (2 -> 9 -> 5).That is, 617 + 295,
// Output:9 -> 1 -> 2,Thatis,912.

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

func (l *LinkedList) SumList(other *LinkedList) *LinkedList {
	carry := 0
	var sum, value1, value2 int
	var tail, head *Node

	l1, l2 := l.Head, other.Head

	for l1 != nil || l2 != nil {
		if l1 != nil {
			value1 = l1.Data
			l1 = l1.Next
		} else {
			value1 = 0
		}
		if l2 != nil {
			value2 = l2.Data
			l2 = l2.Next
		} else {
			value2 = 0
		}

		sum = value1 + value2 + carry
		if sum > 9 {
			carry = 1
			sum -= 10
		} else {
			carry = 0
		}

		node := &Node{Data: sum}
		if tail == nil {
			head = node
			tail = node
		} else {
			tail.Next = node
			tail = node
		}
	}
	if carry == 1 {
		tail.Next = &Node{Data: 1}
	}
	return &LinkedList{Head: head, Tail: tail}
}

func (l *LinkedList) Display() {
	if l.Head == nil {
		fmt.Println("the list is empty")
	}
	curr := l.Head
	for curr != nil {
		fmt.Printf("%v ->", curr.Data)
		curr = curr.Next
	}
	fmt.Println()
}

func main() {
	newList1 := LinkedList{Head: nil, Tail: nil}
	newList1.PushBack(7)
	newList1.PushBack(1)
	newList1.PushBack(6)
	newList1.Display()
	newList2 := LinkedList{Head: nil, Tail: nil}
	newList2.PushBack(5)
	newList2.PushBack(9)
	newList2.PushBack(2)
	newList2.Display()

	sum := newList1.SumList(&newList2)
	sum.Display()
}
