package main

import "fmt"

func main() {
	fmt.Println("Declaration of Variables")

	var firstVar string = "first_variable"
	fmt.Println(firstVar)

	secondVar := "second_variable"
	fmt.Println(secondVar)

	var (
		thirdVar string = "third_variable"
		lastvar string = "last_variable"
	)

	fmt.Println(thirdVar, lastvar)

	newVar, newSecondVar := "new_variable", "other_new_variable"
	fmt.Println(newVar, newSecondVar)

	const firstConst string = "constant_var"
	fmt.Println(firstConst)
}