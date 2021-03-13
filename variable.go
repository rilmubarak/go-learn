package main

import "fmt"

func main() {
	var name string

	name = "Mr Rob"
	fmt.Println(name)

	name = "My Real Name is Robinson!"
	fmt.Println(name)

	var address = "Jl. Pocuk Baren"
	fmt.Println("Alalamat =", address)

	var age = 24
	fmt.Println("Age =", age)

	Status := "Singgle"
	fmt.Println(Status)

	var (
		Country = "Indonesia "
		Flag    = "Red " + "White"
	)

	fmt.Println(Country + Flag)
}
