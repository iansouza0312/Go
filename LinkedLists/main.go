package main

import "fmt"

type Pessoa struct {
	Name     string
	Lastname string
	Age      int8
}

func main() {
	list := List{}

	ian := Pessoa{"Ian", "Souza", 19}
	joao := Pessoa{"Joao", "Santos", 37}
	ana := Pessoa{"Ana", "Almeida", 54}
	maria := Pessoa{"Maria", "Santos", 22}
	jose := Pessoa{"Jose", "Augusto", 16}

	list.Append(ian)
	list.Append(joao)
	list.Append(ana)
	list.Append(maria)
	list.Append(jose)

	list.Display()
	fmt.Println("----------------------------------------------------")

	pessoa := list.Search("Maria")
	fmt.Println(pessoa)
	fmt.Println("----------------------------------------------------")

	list.Remove("Ana")
	list.Display()
}