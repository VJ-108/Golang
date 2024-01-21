package main

import (
	"fmt"
)

func main() {
	fmt.Println("Arrays")

	var fruitList [4]string
	fruitList[0] = "Apple"
	fruitList[1] = "Tomato"
	fruitList[3] = "Peach"

	fmt.Println("Fruit List: ", fruitList)  //It prints blank space if no value is present
	fmt.Println("Length: ", len(fruitList)) //will print actual length and not no. of elements

	var vegList = [5]string{"potato", "beans", "mushroom"}
	fmt.Println("Length: ", len(vegList)) //5

}
