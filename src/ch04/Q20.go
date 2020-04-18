package ch04

import (
	"container/list"
	"errors"
	"fmt"
)

func Q20() {
	fmt.Println("=========Q20=========")
	//The following is the implementation of a program using doubly linked lists from container/list.
	fmt.Println("---t201---")
	t201()
	fmt.Println("---t202---")
	t202()
}

func t201() {
	l := list.New()
	l.PushBack(1)
	l.PushBack(2)
	l.PushBack(4)
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Printf("t201======%v\n", e.Value)
	}
}

type Value int
type Node struct {
	Value
	prev, next *Node
}
type List struct {
	head, tail *Node
}

func (l *List) Front() *Node {
	return l.head
}
func (n *Node) Next() *Node {
	return n.next
}
func (l *List) Push(v Value) *List {
	n := &Node{Value: v}
	if l.head == nil {
		l.head = n
	} else {
		l.tail.next = n
		n.prev = l.tail
	}
	l.tail = n
	return l
}

var errEmpty = errors.New("List is empty")

func (l *List) Pop() (v Value, err error) {
	if l.tail == nil {
		err = errEmpty
	} else {
		v = l.tail.Value
		l.tail = l.tail.prev
		if l.tail == nil {
			l.head = nil
		}
	}
	return v, err
}
func t202() {
	l := new(List)
	l.Push(1)
	l.Push(2)
	l.Push(4)
	for n := l.Front(); n != nil; n = n.Next() {
		fmt.Printf("%v\n", n.Value)
	}
	fmt.Println()
	for v, err := l.Pop(); err == nil; v, err = l.Pop() {
		fmt.Printf("%v\n", v)
	}
}
