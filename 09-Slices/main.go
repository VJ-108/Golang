package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Slices")

	var fruitList = []string{"Apple", "Tomato", "Peach"}
	fmt.Printf("Type of fruitList: %T\n", fruitList)

	fruitList = append(fruitList, "Mango", "Banana")
	fmt.Println(fruitList)

	fruitList = append(fruitList[1:])
	fmt.Println(fruitList)

	highScore := make([]int, 4)

	highScore[0] = 0
	highScore[1] = 1
	highScore[2] = 2
	highScore[3] = 3
	// highScore[4] = 4 gives error

	highScore = append(highScore, 4, 5, 6) //Append automatically reallocates memory
	fmt.Println(highScore)

	fmt.Println(sort.IntsAreSorted(highScore)) //true
	sort.Ints(highScore)
	fmt.Println(highScore)

	//How to remove a value from slices based on index

	var courses = []string{"reactjs", "js", "swift", "python", "ruby"}
	fmt.Println(courses)
	var i int = 2
	courses = append(courses[:i], courses[i+1:]...)
	fmt.Println(courses)
}
