package main

import "fmt"

func main() {
	//var person map[string]string = map[string]string{}
	//person["name"] = "Nurul"
	//person["address"] = "Malang"

	person := map[string]string{
		"name":    "Nurul",
		"address": "Malang",
	}

	fmt.Println(person["name"])
	fmt.Println(person["address"])
	fmt.Println(person)

	book := make(map[string]string)
	book["title"] = "Buku Golang"
	book["author"] = "Nurul"
	book["ups"] = "Salah"

	fmt.Println(book)

	delete(book, "ups")

	fmt.Println(book)
}
