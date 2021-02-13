package main

import "fmt"

func main() {
	type NoKTP string
	type Married bool

	var noKTP NoKTP = "314062706950005"
	var marriedStatus Married = true

	fmt.Println(noKTP)
	fmt.Println(marriedStatus)

}
