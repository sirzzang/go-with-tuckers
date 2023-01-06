package main

import (
	"fmt"
	"sync"
)

func square(wg *sync.WaitGroup, ch chan int) {
	for n := range ch {
		fmt.Printf("square from channel: %d\n", n*n)
	}
	wg.Done()
}

func main() {

	var wg sync.WaitGroup

	// make channel
	ch := make(chan int) // empty channel

	// add square goroutine
	wg.Add(1)
	go square(&wg, ch)

	// send data to channel
	for i := 0; i < 10; i++ {
		ch <- i
	}

	// close channel
	close(ch)

	// wait for goroutine to be done
	wg.Wait()
}
