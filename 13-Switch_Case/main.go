package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Switch Case")
	rand.Seed(time.Now().UnixNano())
	diceNumber := rand.Intn(6) + 1
	fmt.Println("Value of Dice: ", diceNumber)
	switch diceNumber {
	case 1:
		fmt.Println("1")
	case 2:
		fmt.Println("2")
	case 3:
		fmt.Println("3")
		fallthrough //next case will run automatically
	case 4:
		fmt.Println("4")
		fallthrough
	case 5:
		fmt.Println("5")
	case 6:
		fmt.Println("6")
	}
}
