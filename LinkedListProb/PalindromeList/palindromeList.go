//Palindrome: Implement a function to check if a linked list is a palindrome.

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
		fmt.Println("the list is empty")
	}
	curr := l.Head
	for curr != nil {
		fmt.Printf("%v ->", curr.Data)
		curr = curr.Next
	}
	fmt.Println()
}

func (l *LinkedList) PalindromeList() bool {

	//find the middle node of the list
	slow := l.Head
	fast := l.Head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	//reverse the second half of the list
	curr := slow
	var prev *Node
	prev = nil
	for curr != nil {
		temp := curr.Next
		curr.Next = prev
		prev = curr
		curr = temp
	}

	//compare the first half of the list with the second half of the list
	left, right := l.Head, prev
	for right != nil {
		if left.Data != right.Data {
			return false
		}
		left = left.Next
		right = right.Next
	}

	//restore the original list
	curr = prev
	prev = nil
	for curr != nil {
		temp := curr.Next
		curr.Next = prev
		prev = curr
		curr = temp
	}
	l.Tail.Next = prev
	return true
}

func main() {
	newList := LinkedList{Head: nil, Tail: nil}
	newList.PushBack(10)
	newList.PushBack(20)
	newList.PushBack(30)
	newList.PushBack(20)
	newList.PushBack(10)
	newList.Display()
	result := newList.PalindromeList()
	fmt.Println(result)
}
