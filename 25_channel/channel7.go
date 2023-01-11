package main

import (
	"fmt"
	"sync"
	"time"
)

func square(wg *sync.WaitGroup, ch chan int) {

	// 1초가 지나면 channel에 Time 타입 데이터 들어옴
	tick := time.Tick(time.Second)            // <-chan Time
	terminate := time.After(10 * time.Second) // <-chan Time

	// tick -> terminate -> ch
	// select문에 여러 데이터 올 경우에, 어떤 게 실행될지는 랜덤
	// 머신마다 다르게 실행된다
	for {
		select {
		case <-tick:
			fmt.Println("tick")
		case <-terminate:
			fmt.Println("terminate")
			wg.Done() // 작업 완료
		case n := <-ch:
			fmt.Printf("n in channel, square: %d\n", n*n)
			time.Sleep(time.Second)
		}
	}
}

func main() {

	var wg sync.WaitGroup
	ch := make(chan int)

	wg.Add(1)
	go square(&wg, ch)

	for i := 0; i < 10; i++ {
		ch <- i
	}
	wg.Wait()

}
