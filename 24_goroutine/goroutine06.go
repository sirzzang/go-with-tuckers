package main

import (
	"fmt"
	"sync"
	"time"
)

// waitgroup 사용할 때는 sub goroutine이 모두 끝난 후에 waitgroup에 add해서는 안 된다
var wg sync.WaitGroup

func PrintRune() {
	runes := []rune{'가', '나', '다', '라', '마', 'a', 'b', 'A'}
	for _, r := range runes {
		fmt.Printf("%c ", r)
		time.Sleep(300 * time.Millisecond) // sub 고루틴의 흐름을 조금 지연시켜 볼 수 있음
	}

	wg.Done()
}

func PrintNum() {
	for i := 1; i <= 10; i++ {
		fmt.Printf("%d ", i)
		time.Sleep(400 * time.Millisecond)

	}

	wg.Done()
}

// time.Sleep을 통해 wait이 없더라도 main 고루틴 흐름이 멈춰지게 됨
func main() {
	go PrintRune()
	go PrintNum()

	time.Sleep(3 * time.Second) // main 고루틴이 계속 진행 중
	wg.Wait()                   // main 고루틴 흐름을 멈춰 둠
	wg.Add(2)                   // sub goroutine이 안 끝나는 시점까지는 괜찮지만, 끝나자마자 panic
}
