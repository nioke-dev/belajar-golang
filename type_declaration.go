package main

import "fmt"

func main() {

	type NoKtp string

	var ktpNurul NoKtp = "111111111111"

	var contoh string = "2222222222"
	var contohKtp NoKtp = NoKtp(contoh)

	fmt.Println(ktpNurul)
	fmt.Println(contohKtp)
}
