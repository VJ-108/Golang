package main

import "fmt"

func main() {
	fmt.Println("Function")
	//Can't declare function inside function
	greet()

	result := add(1, 2)
	fmt.Println("Result: ", result)

	Result,_ := adder(1,2,3,4,5) //Returns two things one is int and other is string
	fmt.Println("Result: ",Result)
}

func greet() {
	fmt.Println("Hello")
}

// int written outside brackets tells about return type
func add(a int, b int) int {
	return a + b
}

func adder(values ...int) (int,string) {
	total := 0
	for _, v := range values {
		total += v
	}
	return total,"Hello"
}
