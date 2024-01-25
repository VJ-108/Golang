package main

import "fmt"

func main() {
	fmt.Println("Loops")

	day := []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}

	// fmt.Println(day)

	for i := 0; i < len(day); i++ {
		fmt.Println(day[i])
	}

	for i := range day {
		fmt.Println(day[i])
	}
	for i, v := range day {
		fmt.Printf("Index: %v and value: %v\n", i, v)
	}

	rougueValue := 1
	for rougueValue < 10 {
		if rougueValue == 2 {
			goto lco
		}
		if rougueValue == 5 {
			rougueValue++
			continue
		}
		if rougueValue == 6 {
			break
		}
		fmt.Println("Value: ", rougueValue)
		rougueValue++
	}

lco:
	fmt.Println("Jump")
}
