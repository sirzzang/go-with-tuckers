package main

import "fmt"

func main() {

	a := [5]int{1, 2, 3, 4, 5}
	b := [5]float64{10, 20, 30, 40, 50}

	for i, v := range a {
		fmt.Printf("a[%d] = %d\n", i, v)
	}
	fmt.Println()

	for i, v := range b {
		fmt.Printf("b[%d] = %d\n", i, v)
	}
	fmt.Println()

	b = a // type mismatch
	for i, v := range b {
		fmt.Printf("b[%d] = %d\n", i, v)
	}

}
