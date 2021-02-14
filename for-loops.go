package main

import "fmt"

func main() {

	//For Short Statement
	for i := 0; i < 10; i++ {
		fmt.Println(i, "For Short Statement")
	}

	//Using Slice
	slice := [...]string{"Yana", "Andika", "Mantap"}
	for i := 0; i < len(slice); i++ {
		fmt.Println(slice[i])
	}

	//Using Hashmap
	person := make(map[string]string)
	person["Name"] = "Yana"
	person["Age"] = "25"

	for key, value := range person {
		fmt.Println("Key", key, "Value", value)
		if value == "Yana" {
			fmt.Println("Terdapat Sebuah Value bernama ", value)
		}
	}

	//If You dont want variable key
	for _, value := range person {
		fmt.Println("Value", value)
		if value == "Yana" {
			fmt.Println("Terdapat Sebuah Value bernama ", value)
		}
	}

}
