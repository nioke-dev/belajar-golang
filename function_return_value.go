package main

import "fmt"

func getHello(name string) string {
	hello := "Hello " + name
	return hello
}

func main() {
	result := getHello("Nurul")
	fmt.Println(result)

	fmt.Println(getHello("Budi"))
	fmt.Println(getHello("Joko"))
}
