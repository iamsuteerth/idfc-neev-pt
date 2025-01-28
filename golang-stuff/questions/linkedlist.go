package questions

// interface{} is used to supply function with any value
// Unnecessary complicated shit but itiswhatitis

import (
	"errors"
	"fmt"
)

// Define List and Node types here.
type Node struct {
	Value interface{}
	prev  *Node
	next  *Node
}
type List struct {
	head *Node
	tail *Node
}

func (node Node) String() string {
	return fmt.Sprintf("%s", node.Value)
}
func (dll List) String() string {
	if dll.isEmpty() {
		return "Empty List"
	}
	curr := dll.head
	listToString := curr.String()
	for curr := dll.head.next; curr != nil; curr = curr.next {
		listToString += " - " + curr.String()
	}
	return listToString
}
func NewList(elements ...interface{}) *List {
	dll := &List{}
	for _, v := range elements {
		dll.Push(v)
	}
	return dll
}
func (dll *List) isEmpty() bool {
	return dll.head == nil
}
func (n *Node) Next() *Node {
	return n.next
}
func (n *Node) Prev() *Node {
	return n.prev
}
func (dll *List) Shift() (interface{}, error) {
	if dll.isEmpty() {
		return int(0), errors.New("err empty list")
	}
	headVal := dll.head.Value
	dll.head = dll.head.next
	// After shifting, the head is NILL
	if dll.head == nil {
		dll.tail = nil
	} else {
		dll.head.prev = nil
	}
	return headVal, nil
}
func (dll *List) Unshift(v interface{}) {
	node := &Node{Value: v}
	if dll.isEmpty() {
		dll.head = node
		dll.tail = node
		return
	}
	node.next = dll.head
	dll.head.prev = node
	dll.head = node
}
func (dll *List) Push(v interface{}) {
	node := &Node{Value: v}
	if dll.isEmpty() {
		dll.head = node
		dll.tail = node
		return
	}
	node.prev = dll.tail
	dll.tail.next = node
	dll.tail = node
}
func (dll *List) Pop() (interface{}, error) {
	if dll.isEmpty() {
		return int(0), errors.New("err empty list")
	}
	tailVal := dll.tail.Value
	dll.tail = dll.tail.prev
	if dll.tail == nil {
		dll.head = nil
	} else {
		dll.tail.next = nil
	}
	return tailVal, nil
}
func (dll *List) Reverse() {
	for curr := dll.head; curr != nil; curr = curr.prev {
		curr.prev, curr.next = curr.next, curr.prev
	}
	dll.head, dll.tail = dll.tail, dll.head
	return
}
func (dll *List) First() *Node {
	return dll.head
}
func (dll *List) Last() *Node {
	return dll.tail
}
