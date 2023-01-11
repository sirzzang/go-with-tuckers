package main

import "fmt"

func main() {

	ch := make(chan int) // 크기가 0인 채널

	ch <- 1
	fmt.Println("Never prints...")

}
