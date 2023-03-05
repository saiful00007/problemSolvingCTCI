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

}

func main() {
	newList := LinkedList{Head: nil, Tail: nil}
	newList.PushBack(10)
	newList.PushBack(20)
	newList.PushBack(30)
	newList.PushBack(40)
	newList.PushBack(50)
	newList.Display()
}
