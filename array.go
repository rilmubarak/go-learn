package main

import "fmt"

func main() {
	var names [3]string

	names[0] = "Riki"
	names[1] = "Niki"
	names[2] = "Adi"

	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])

	var values = [3]int{
		90,
		20,
		30,
	}
	fmt.Println(values)
	fmt.Println(values[2])

	fmt.Println(len(names))
	fmt.Println(len(values))
}
