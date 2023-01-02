package main

import "fmt"

type opFunc func(int, int) int

func getOperator(op string) opFunc {
	if op == "+" {
		return func(a, b int) int {
			return a + b
		}
	} else if op == "-" {
		return func(a, b int) int {
			return a - b
		}
	} else if op == "*" {
		return func(a, b int) int {
			return a * b
		}
	} else {
		return nil
	}
}

func main() {
	fn := getOperator("+")
	fmt.Println(fn(3, 4))

	fnNil := getOperator(";") // panic
	fmt.Println(fnNil(3, 4))

}
