package main

import "fmt"

func main() {
	fmt.Println("Pointers")
	var ptr1 *int
	fmt.Println("Default value of pointer: ", ptr1)

	num := 3
	var ptr = &num
	fmt.Println("Value of pointer: ", ptr)
	fmt.Println("Value of num: ", *ptr)

	*ptr = *ptr * 2
	fmt.Println("Updated value of num: ", num)
}
