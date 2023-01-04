package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func SumAtoB(a, b int) {
	sum := 0
	for i := a; i < b; i++ {
		sum += i
	}
	fmt.Printf("from %d to %d, sum is %d\n", a, b, sum)
	wg.Done()
}

func main() {
	wg.Add(10)
	for i := 0; i < 14; i++ { // TODO: goroutine 개수 다르면 오류가 나지는 않고 아예 실행이 안 되거나 끝나거나?
		go SumAtoB(1, 1000000000)
	}
	wg.Wait()
	fmt.Println("main goroutine bye")
}
