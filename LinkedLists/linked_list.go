package main

import "fmt"

type List struct {
	Head *Node
	Tail *Node
}

func (l *List) Append(pessoa Pessoa) {
	node := &Node{Value: pessoa}
	if l.Head == nil {
		l.Head = node
	}
	if l.Tail != nil {
		l.Tail.Next = node
	}
	l.Tail = node
}

func (l *List) Display() {
	node := l.Head
	for node != nil {
		fmt.Println(node.Value.Name)
		node = node.Next
	}
}

type Node struct {
	Value Pessoa
	Next  *Node
}