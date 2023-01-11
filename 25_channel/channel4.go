package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	// 크기가 있는 채널
	ch := make(chan int, 2)

	wg.Add(1)
	go square(&wg, ch)

	// 채널에 데이터 전송
	for i := 0; i < 10; i++ {
		ch <- i
	}

	// square 고루틴 종료 대기
	wg.Wait()

}

// 좀비 고루틴
func square(wg *sync.WaitGroup, ch chan int) {

	// 버퍼가 있는 경우도 마찬가지로 계속 대기
	for n := range ch {
		fmt.Printf("square from channel: %d\n", n*n)
	}

	wg.Done()

}
