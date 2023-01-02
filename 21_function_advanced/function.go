package main

import "fmt"

func sum(nums ...int) int {
	sum := 0

	// 가변 인수는 함수 내부에서 해당 타입의 슬라이스로 처리
	for _, v := range nums {
		sum += v
	}
	return sum
}

func main() {
	fmt.Println(sum(1, 2, 3, 4, 5))
	// fmt.Println(sum(10, 20, 3.14, "string"))
}
