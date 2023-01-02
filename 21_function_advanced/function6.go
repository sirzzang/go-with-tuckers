package main

import "fmt"

func main() {

	// 함수 리터럴
	fn := func(a, b int) int {
		return a + b
	}
	fmt.Println(fn(3, 4))

	// 함수 리터럴 바로 호출
	fmt.Println(func(a, b int) int {
		return a * b
	}(3, 4))
}
