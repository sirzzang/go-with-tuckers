package main

import "fmt"

var cnt int = 0

func IncreaseAndReturn() int {
	fmt.Println("Called", cnt)
	cnt++
	return cnt
}

func main() {
	// short circuit
	if false && IncreaseAndReturn() < 5 {
		fmt.Println("here")
	}

	// short circuit
	if true || IncreaseAndReturn() > 10 {
		fmt.Println("here2")
	}

	// IncreaseAndReturn never called
	fmt.Println(cnt)
}
