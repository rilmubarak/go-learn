package main

import "fmt"

func main() {
	var Ujian = 89
	var Absensi = 80

	var LulusUjian = Ujian > 80
	var LulusAbsensi = Absensi > 80

	var lulus = LulusUjian && LulusAbsensi

	if lulus == true {
		fmt.Println("lulus")
	} else {
		fmt.Println("Tidak Lulus")
	}

	fmt.Println(lulus)
}
