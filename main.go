package main

import "fmt"

func main() {
	fmt.Println("Who are you?")

	var input string
	fmt.Scanln(&input)

	var hello string = "Hello " + input + ", this is my journey to understanding Go programming language fundamentaly, just in case you want to know."

	fmt.Println(hello)
}
