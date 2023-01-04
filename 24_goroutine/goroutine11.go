package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func SumAtoB(a, b int) {
	sum := 0
	for i := a; i < b; i++ {
		sum += i
	}
	fmt.Printf("from %d to %d, sum is %d\n", a, b, sum)
	wg.Done()
}

func main() {
	for i := 0; i < 10; i++ {
		go SumAtoB(1, 2)
	}
	wg.Add(10)
	wg.Wait()
	// fmt.Println("main goroutine bye")
}
