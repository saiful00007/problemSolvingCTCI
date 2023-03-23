//Impliment a MyQueue class which implements a queue using tow stacks..

package main

import "fmt"

type MyQueue struct {
	stack1 []int
	stack2 []int
}

func (q *MyQueue) Enqueue(value int) {
	q.stack1 = append(q.stack1, value)
}

func (q *MyQueue) Dequeue() int {
	if len(q.stack2) == 0 {
		for len(q.stack1) > 0 {
			n := len(q.stack1) - 1
			q.stack2 = append(q.stack2, q.stack1[n])
			q.stack1 = q.stack1[:n]
		}
	}
	n := len(q.stack2) - 1
	x := q.stack2[n]
	q.stack2 = q.stack2[:n]
	return x
}

func (q *MyQueue) Peek() int {
	if len(q.stack2) == 0 {
		for len(q.stack1) > 0 {
			n := len(q.stack1) - 1
			q.stack2 = append(q.stack2, q.stack1[n])
			q.stack1 = q.stack1[:n]
		}
	}
	return q.stack2[len(q.stack2)-1]
}

func (q *MyQueue) IsEmpty() bool {
	return len(q.stack1) == 0 && len(q.stack2) == 0
}

func main() {
	newQueue := MyQueue{}

	newQueue.Enqueue(1)
	newQueue.Enqueue(2)
	newQueue.Enqueue(3)
	newQueue.Enqueue(4)

	for !newQueue.IsEmpty() {
		fmt.Println(newQueue.Dequeue())
	}
}
