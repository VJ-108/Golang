package main

import (
	"fmt"
	"net/http"
	"sync"
)

var signals = []string{"test"}

var wg sync.WaitGroup //pointers
var mut sync.Mutex    //pointer
func main() {
	fmt.Println("Go routines")
	// go greeter("Hello")
	// greeter("World")
	websiteList := []string{
		"https://google.com",
		"https://go.dev",
		"https://fb.com",
	}
	for _, web := range websiteList {
		go getStatusCode(web)
		wg.Add(1)
	}
	wg.Wait()
	fmt.Println(signals)
}

// func greeter(s string) {
// 	for i := 0; i < 5; i++ {
// 		fmt.Println(s)
// 	}
// }

func getStatusCode(endpoint string) {
	defer wg.Done()
	res, err := http.Get(endpoint)
	if err != nil {
		panic(err)
	}
	mut.Lock()
	signals = append(signals, endpoint)
	mut.Unlock()
	fmt.Printf("%d status Code for %s\n", res.StatusCode, endpoint)
}
