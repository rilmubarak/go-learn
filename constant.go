package main

import "fmt"

func main() {
	fmt.Println("Belajar Constanta")

	const name = "Baril"
	fmt.Println(name)

	const (
		firstName = "Baril "
		lastName  = "M"
	)

	fmt.Println(firstName + lastName)
}
