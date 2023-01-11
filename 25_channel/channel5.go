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

	// 채널에 데이터전송
	for i := 0; i < 10; i++ {
		ch <- i
	}

	// 데이터 전송을 마치면 채널을 닫음
	close(ch)

	// square 고루틴 종료 대기
	wg.Wait()
}

func square(wg *sync.WaitGroup, ch chan int) {

	// 채널이 비어 있는 상태에서 닫히게 되면 for range 문 빠져 나감
	for n := range ch {
		fmt.Printf("square from channel: %d\n", n*n)
	}

	// 종료
	wg.Done()
}
