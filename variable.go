package main

import "fmt"

func main() {
	//Variable Apa bila tidak digunakan menjadi error

	// Deklarasi Multiple variable
	var (
		firstName = "Yana"
		lastName  = "Andika"
	)

	fmt.Println("First Name ", firstName)
	fmt.Println("Last Name ", lastName)

	//Langsung Set Data
	name := "Yana Andika"

	fmt.Println(name)

	name = "Yana Bedugul"
	fmt.Println(name)

	var age = 38
	fmt.Println(age)
}
