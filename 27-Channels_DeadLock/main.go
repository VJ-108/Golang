package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Channels and DeadLock")

	myCh := make(chan int, 5)
	wg := &sync.WaitGroup{}
	// myCh <- 5
	// fmt.Println(<-myCh)
	wg.Add(2)
	//recieve only channel
	go func(ch <-chan int, wg *sync.WaitGroup) {
		val, isChanelOpen := <-myCh
		fmt.Println(isChanelOpen)
		fmt.Println(val)
		// fmt.Println(<-myCh)
		wg.Done()
	}(myCh, wg)
	//send only channel
	go func(ch chan<- int, wg *sync.WaitGroup) {
		myCh <- 0
		// myCh <- 6
		close(myCh)
		wg.Done()
	}(myCh, wg)
	wg.Wait()
}
