package main

import "fmt"

type OpFunc func(int, int) int

func Process(a, b int, op OpFunc) {
	fmt.Println(op(a, b))
}

func main() {
	op := func(a, b int) int {
		return a * b
	}
	Process(5, 6, op) // 30
}
