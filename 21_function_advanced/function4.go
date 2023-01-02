package main

import "fmt"

func add(a, b int) int {
	return a + b
}

func mul(a, b int) int {
	return a * b
}

func getOperator(op string) func(a, b int) int {
	if op == "+" {
		return add
	} else if op == "*" {
		return mul
	} else {
		return nil // 유효하지 않은 함수
	}
}

func main() {

	var operator func(int, int) int // 함수 타입 변수
	operator = getOperator("+")     // 함수 타입 변수에 함수 타입 값 대입

	var result int
	result = operator(3, 4)
	fmt.Println(result)
}
