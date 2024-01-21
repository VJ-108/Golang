package main

import "fmt"

//Can't use walrus operator globally

const hello = "Hello"

func main() {
	var username string = "Hello"
	fmt.Println(username)
	fmt.Printf("Variable is of type: %T\n", username)

	var isLoggedIn bool = true
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is of type: %T\n", isLoggedIn)

	var smallint uint8 = 255
	fmt.Println(smallint)
	fmt.Printf("Variable is of type: %T\n", smallint)

	var smallfloat float32 = 254.823727328
	fmt.Println(smallfloat)
	fmt.Printf("Variable is of type: %T\n", smallfloat)

	//int,float has default value as 0, while bool has false

	//no var style, ':=' - walrus operator

	num := 3000
	fmt.Println(num)
	fmt.Printf("Variable is of type: %T\n", num)

	fmt.Println(hello)
}
