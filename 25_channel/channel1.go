package main

import (
	"fmt"
	"time"
)

func main() {
	// ch := make(chan int)
	ch := make(chan int, 2)

	go square()
	ch <- 1 // 크기가 0인 채널의 경우, 데이터를 빼 가기를 무한히 대기
	fmt.Println("Never prints when channel length is 0")
}

func square() {
	for {
		time.Sleep(time.Second)
		fmt.Println("sleep")
	}
}
