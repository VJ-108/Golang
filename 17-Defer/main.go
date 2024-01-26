package main

import "fmt"

func main() {
	defer fmt.Println("Hello 1") //defer puts this line at the last of main so it will execute at last
	defer fmt.Println("Hello 2") // defer put them just like a stack
	defer fmt.Println("Hello 3")
	fmt.Println("Defer")
	myDefer()
}

func myDefer() {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}
