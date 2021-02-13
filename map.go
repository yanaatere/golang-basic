package main

import "fmt"

func main() {
	person := map[string]string{
		"name":    "Yana",
		"address": "Jakarta",
	}

	fmt.Println(person)
	fmt.Println(person["name"])

	//How To Add New Key
	person["title"] = "Programmer"

	fmt.Println(person)
	fmt.Println(person["title"])

	//Function Make New Map
	book := make(map[string]string)
	book["title"] = "Buku GO-Lang"
	book["author"] = "Yana"
	book["wrong"] = "ups"
	fmt.Println(len(book))

	delete(book, "wrong")
	fmt.Println(book)
}
