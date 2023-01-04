package main

import (
	"fmt"
	"time"
)

func PrintRune() {
	runes := []rune{'가', '나', '다', '라', '마', 'a', 'b', 'A'}
	for _, v := range runes {
		time.Sleep(300 * time.Millisecond)
		fmt.Printf("%c ", v)
	}
}

func PrintNum() {
	for i := 1; i < 10; i++ {
		time.Sleep(400 * time.Millisecond)
		fmt.Printf("%d ", i)
	}
}

func main() {
	go PrintNum()
	go PrintRune()
	time.Sleep(3 * time.Second)
}
