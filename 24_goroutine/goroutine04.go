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
	}

	wg.Done()
}

func PrintNum() {
	for i := 1; i <= 10; i++ {
		fmt.Printf("%d ", i)
	}

	wg.Done()
}

// time.Sleep을 통해 wait이 없더라도 main 고루틴 흐름이 멈춰지게 됨
func main() {
	go PrintRune()
	go PrintNum()

	time.Sleep(3 * time.Second) // main 고루틴이 진행 중
	wg.Add(2)                   // main 고루틴이 안 끝난 시점에 이미 goroutine이 모두 종료된 후이므로 add할 고루틴이 없어 panic
}
