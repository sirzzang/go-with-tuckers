package main

import (
	"fmt"
	"sync"
)

func square(wg *sync.WaitGroup, ch chan int, quit chan bool) {

	for {
		select {
		// if data in channel, then square
		case n := <-ch:
			fmt.Printf("square from channel: %d\n", n*n)
		// if to be quitted, return
		case <-quit:
			fmt.Println("bye")
			wg.Done()
			return
		}
	}
}

func main() {

	var wg sync.WaitGroup

	// make channel
	ch := make(chan int)
	quit := make(chan bool)

	// add square goroutine
	wg.Add(1)
	go square(&wg, ch, quit)

	// send data to channel
	for i := 0; i < 10; i++ {
		ch <- i
	}

	// notify to quit
	quit <- true

	// wait for goroutine to be done
	wg.Wait()
}
