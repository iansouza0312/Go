package main

import (
	"adressTest/adress"
	"fmt"
)

func main() {
	adressType := adress.AdressType("Rua 12 de março")
	fmt.Println(adressType)
}