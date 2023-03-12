// Intersection; Given two (singly) linked lists, determine if the two lists intersect. Return the intersecting node. Note that the intersection is defined based on reference, not value. That is, if the kth
// node of the first linked list is the exact same node (by reference) as the jt h node of the second
// linked list, then they are intersecting.

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

func (l *LinkedList) Len() int {
	count := 0
	curr := l.Head
	for curr != nil {
		count++
		curr = curr.Next
	}
	return count
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
	curr := l.Head
	if curr == nil {
		fmt.Println("the list is empty")
	}
	for curr != nil {
		fmt.Printf("%v ->", curr.Data)
		curr = curr.Next
	}
	fmt.Println()
}

func (l *LinkedList) IntersectionNode(other *LinkedList) *Node {
	if l.Head == nil || other.Head == nil {
		return nil
	}
	//find the length of the two list
	lenL := l.Len()
	lenOther := other.Len()

	//move the longer linked list pointer ahead by the difference in length
	ptr1 := l.Head
	ptr2 := other.Head

	if lenL > lenOther {
		for i := 0; i < lenL-lenOther; i++ {
			ptr1 = ptr1.Next
		}
	} else if lenOther > lenL {
		for i := 0; i < lenOther-lenL; i++ {
			ptr2 = ptr2.Next
		}
	}
	//now traverse both linked list together and return the intersection node
	for ptr1 != nil && ptr2 != nil {
		if ptr1 == ptr2 {
			return ptr1
		}
		ptr1 = ptr1.Next
		ptr2 = ptr2.Next
	}
	return nil
}

func main() {
	list1 := &LinkedList{Head: nil, Tail: nil}
	list1.PushBack(4)
	list1.PushBack(1)
	list1.PushBack(8)
	list1.PushBack(4)
	list1.PushBack(5)
	list1.Display()

	list2 := &LinkedList{Head: nil, Tail: nil}
	list2.PushBack(5)
	list2.PushBack(6)
	list2.PushBack(1)
	list2.PushBack(8)
	list2.PushBack(4)
	list2.PushBack(5)
	list2.Display()

	intersectionNode := list2.IntersectionNode(list1)
	if intersectionNode != nil {
		fmt.Println("Intersection Node: ", intersectionNode.Data)
	} else {
		fmt.Println("no intersection found")
	}
}
