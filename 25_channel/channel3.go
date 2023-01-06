package main

import (
	"fmt"
	"sync"
)

func square(wg *sync.WaitGroup, ch chan int) {

	// square until channel is empty
	for n := range ch {
		fmt.Println(ch)
		fmt.Printf("square from channel: %d\n", n*n)
	}

	// notify work done
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

	// wait for goroutine to be done
	wg.Wait()

}
