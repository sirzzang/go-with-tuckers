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
		// time.Sleep(300 * time.Millisecond)
	}

	wg.Done()
}

func PrintNum() {
	for i := 1; i <= 10; i++ {
		fmt.Printf("%d ", i)
		// time.Sleep(400 * time.Millisecond)
	}

	wg.Done()
}

func main() {

	wg.Add(2)

	go PrintRune()
	go PrintNum()

	wg.Wait() // main 고루틴 흐름 블락. main 고루틴에서 sleep하지 않아도 됨

}
