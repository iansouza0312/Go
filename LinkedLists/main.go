package main

type Pessoa struct {
	Name     string
	Lastname string
	Age      int8
}

func main() {
	list := List{}

	ian := Pessoa{"Ian", "Souza", 19}
	joao := Pessoa{"Joao", "Santos", 37}

	list.Append(ian)
	list.Append(joao)

	list.Display()
}