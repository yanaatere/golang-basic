package main

import "fmt"

func main() {
	var names [3]string

	names[0] = "Yana"
	names[1] = "Andika"

	fmt.Println(names)
	fmt.Println(names[1])

	//Deklarasi Multiple Array
	var values = [3]int{
		90, 30, 50,
	}

	fmt.Println("Multiple Array", values)

	var panjangArray [10]string

	// Function array
	fmt.Println("Untuk Menampilkan Panjang Array", len(values))
	fmt.Println("Untuk Menampilkan Panjang Array", len(panjangArray))
}
