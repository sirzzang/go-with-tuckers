package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func PrintRune() {
	runes := []rune{'가', '나', '다', '라', '마', 'a', 'b', 'A', 'ㅋ', 'ㅎ', ';'}
	for _, r := range runes {
		fmt.Printf("%c ", r)
		// time.Sleep(100 * time.Millisecond) // sub 고루틴의 실행이 지연되지 않음
	}

	wg.Done()
}

func PrintNum() {
	for i := 1; i <= 10; i++ {
		fmt.Printf("%d ", i)
		// time.Sleep(300 * time.Millisecond)
	}

	wg.Done()
}

// 애초에 main 고루틴이 먼저 끝나기 때문에 add 하나 안 하나 panic 일어나지 않음
func main() {

	go PrintRune()
	go PrintNum()

	wg.Add(2) // panic 일어나지 않음

}
