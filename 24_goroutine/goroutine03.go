package main

import (
	"fmt"
	"sync"
)

// waitgroup 사용할 때는 sub goroutine이 모두 끝난 후에 waitgroup에 add해서는 안 된다
// 아래 코드는 뒤의 것들과는 달리 main 고루틴도 바로 종료돼서 panic이 나지는 않으나,
// 이렇게 쓴다고 해도 의미가 없다.
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

// wait을 하더라도, main 고루틴에서 멈춰둘 흐름이 없으므로 main 고루틴 바로 종료
func main() {

	go PrintRune()
	go PrintNum()

	wg.Wait() // 진행 중인 main 고루틴이 없음
	wg.Add(2) // 이미 main 고루틴이 종료된 후로, panic 발생하지 않음

}
