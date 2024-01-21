package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Time Handling")
	presentTime := time.Now()
	fmt.Println(presentTime)
	fmt.Println(presentTime.Format("01-02-2006 15:04:05 Monday"))

	createDate := time.Date(2024, time.January, 10, 23, 23, 0, 0, time.UTC)
	fmt.Println(createDate)
	fmt.Println(createDate.Format("01-02-2006 15:04:05 Monday"))


	//Can build exe files for different os using
	//windows - GOOS="windows" go build
	//linux - GOOS="linux" go build
}
