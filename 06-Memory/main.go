package main

import "fmt"

func main() {
	fmt.Println("Memory Management")
	// 	Memory Management:
	// - new():
	//     - Allocate memory but no INIT
	//     - Will get a memory address
	//     - zeroed Storage (can't put data initially)
	// - make():
	//     - Allocate memory and INIT
	//     - Will get a memory address
	//     - non-zeroed Storage (can put data initially)

	// - Garbage Collection happens automatically (if out of scope or nil)

}
