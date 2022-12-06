package main

import "fmt"

func PrintNo(n int) {

	// exit condition
	if n == 0 {
		return
	}
	fmt.Println(n)
	PrintNo(n - 1) // recursive call
	fmt.Println("After recursive call: n = ", n)
}

func main() {
	PrintNo(5)
}
