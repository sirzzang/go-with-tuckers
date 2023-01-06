package main

import "fmt"

func main() {

	ch := make(chan int) // 빼갈 때까지 대기

	ch <- 1
	fmt.Println("Never prints...")

}
