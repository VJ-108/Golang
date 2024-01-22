package main

import "fmt"

func main() {
	fmt.Println("Structs")
	//No inheritance in golang; No super or parent

	hello := User{"Hello", "Hello@gmail.com", true, 16}
	fmt.Println(hello)

	fmt.Printf("Hello details are: %+v\n", hello)
	fmt.Printf("Name: %v\n", hello.Name)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
