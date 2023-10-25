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

func (l *List) Search(nome string) (pessoa Pessoa){
	node := l.Head
	for node != nil{
		if node.Value.Name == nome{
			pessoa = node.Value
			break
		}
		node = node.Next
	}

	return
}

func (l *List) Remove(nome string){
	if l.Head.Value.Name == nome{
		l.Head = l.Head.Next
		return
	}

	prev := l.Head
	node := l.Head.Next

	for node != nil{
		if node.Value.Name == nome{
			prev.Next = node.Next
			break
		}
		prev = node
		node = node.Next
	}

	if l.Tail == node{
		l.Tail = prev
	}
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