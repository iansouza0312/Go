package main

import "fmt"

func main() {
	stack := Stack{Size: 4}

	stack.Push("Ian")
	stack.Push("Thiago")
	stack.Push("Ana")
	stack.Push("Joao")

	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())
}