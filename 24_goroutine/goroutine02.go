package main

import (
	"fmt"
	"sync"
)

// waitgroup 생성 후 add한 뒤 wait하지 않으면 의미가 없음
var wg sync.WaitGroup

func PrintRune() {
	runes := []rune{'가', '나', '다', '라', '마', 'a', 'b', 'A', 'ㅋ', 'ㅎ', ';'}
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

// main 고루틴에서 실행할 흐름이 없으므로 main 고루틴은 바로 종료
func main() {

	wg.Add(2) // main 고루틴이 바로 종룔되므로, 바로 종료

	go PrintRune()
	go PrintNum()

	// wg.Add(2) // 여기 있어도 마찬가지

}
