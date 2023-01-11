package main

import (
	"fmt"
	"sync"
)

func main() {

	var wg sync.WaitGroup

	ch := make(chan int)
	quit := make(chan bool)

	wg.Add(1)
	go square(&wg, ch, quit)

	// 숫자를 10번에 걸쳐 ch에 넣은 뒤, quit 채널에 데이터 넣음
	for i := 0; i < 10; i++ {
		ch <- i
	}
	quit <- true

	wg.Wait()
}

func square(wg *sync.WaitGroup, ch chan int, quit chan bool) {

	// ch, quit 채널을 무한히 대기하면서 select문 실행
	for {
		select {
		// ch에 오면 여기 실행
		case n := <-ch:
			fmt.Printf("square from channel: %d\n", n*n)
		// quit이 오면 여기 실행
		case <-quit:
			fmt.Println("bye")
			wg.Done()
			return
		}
	}

}
