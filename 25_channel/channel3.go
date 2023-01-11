package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	ch := make(chan int)

	wg.Add(1)
	go square(&wg, ch)

	// 채널에 데이터 전송
	for i := 0; i < 10; i++ {
		ch <- i
	}
	wg.Wait() // main 고루틴은 square 고루틴이 종료되기까지 기다림

}

// 좀비 고루틴
func square(wg *sync.WaitGroup, ch chan int) {

	// for range 문에서 채널이 비어 있더라도 데이터가 계속 들어오기를 기다림
	for n := range ch {
		fmt.Printf("square from channel: %d\n", n*n)
	}

	// 끝나지 않음
	wg.Done()

}
