package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter a number: ")

	//comma ok || err err syntax
	// input, err := reader.ReadString('\n')
	input, _ := reader.ReadString('\n')
	fmt.Println("Number: ", input)

	fmt.Printf("Number is of Type: %T\n", input)
}
