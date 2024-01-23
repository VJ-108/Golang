package main

import "fmt"

func main() {
	fmt.Println("If Else")

	login := 10

	var result string

	if login > 10 {
		result = "Regular User"
	} else if login < 10 {
		result = "WatchOut"
	} else {
		result = "OK"
	}
	fmt.Println(result)

	if 9%2==0{
		fmt.Println("Number is Even")
	}else{
		fmt.Println("Number is Odd")
	}

	if num:=3; num<10{
		fmt.Println("num is less than 10")
	}else{
		fmt.Println("num is greater than or equal to 10")
	}
}
