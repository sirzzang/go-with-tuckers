package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	var wg sync.WaitGroup

	ch := make(chan int) // 크기가 0인 채널

	wg.Add(1)
	go square(&wg, ch)

	ch <- 9 // goroutine에 데이터 전달
	wg.Wait()

}

// waitgroup을 넘겨 받아야(참조로 넘겨 받아야) goroutine에서 끝났다고 알려줄 수 있음
func square(wg *sync.WaitGroup, ch chan int) {
	n := <-ch
	time.Sleep(100 * time.Millisecond)
	fmt.Printf("square from ch: %d\n", n*n)
	wg.Done()
}
