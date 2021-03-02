package main

import "fmt"

func main() {
	firstName, middleName := getFullName()
	fmt.Println(firstName, middleName)
}

func getFullName() (string, string) {
	return "Yana", "Andika"
}
