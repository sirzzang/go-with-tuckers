package main

import "fmt"

type intOpFunc func(int, int) int

func add(a, b int) int {
	return a + b
}

func mul(a, b int) int {
	return a * b
}

func getOperator(op string) intOpFunc {
	if op == "+" {
		return add
	} else if op == "*" {
		return mul
	} else {
		return nil
	}
}

func main() {
	var operator intOpFunc
	operator = getOperator("*")
	result := operator(3, 4)
	fmt.Println(result)
}
