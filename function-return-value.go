package main

import "fmt"

func main() {
	result := sayHello("Yana")
	fmt.Println(result)
}

func sayHello(name string) string {
	if name == "" {
		return "hello"
	} else {
		return "Hello " + name
	}

}
