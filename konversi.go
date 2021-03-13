package main

import "fmt"

func main() {
	var number32 int32 = 32132
	var number64 = int64(number32)

	fmt.Println(number64)

	var name = "Baril"
	var B = name[2]
	var BString string = string(B)

	fmt.Println(BString)
}
