package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
	// "math/rand"
)

func main() {
	fmt.Println("Math")

	//random number
	// rand.Seed(time.Now().UnixNano())
	// fmt.Println(rand.Intn(5)+1)

	//random from crypto
	random, _ := rand.Int(rand.Reader, big.NewInt(5))
	fmt.Println(random)
}
