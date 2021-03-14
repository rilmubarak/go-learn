package main

import "fmt"

func main() {
	type NoKtp string
	type Marriage bool

	var noKtpAsli NoKtp = "123242304323424"
	var MarriageStatus Marriage = true
	fmt.Println(noKtpAsli)
	fmt.Println(MarriageStatus)
}
