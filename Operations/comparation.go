package main

import "fmt"

func main() {
	var name1 = "Baril"
	var name2 = "Baril"
	var name3 = "M"

	var result = name1 == name2
	var result2 = name1 == name3
	var result3 = name1 > name3
	var result4 = name1 < name3
	fmt.Println(result)
	fmt.Println(result2)
	fmt.Println(result3)
	fmt.Println(result4)
}
