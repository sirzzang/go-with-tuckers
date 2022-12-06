package main

import "fmt"

func main() {

	const PI1 float64 = 3.1415926535
	var PI2 float64 = 3.14

	// PI1 = 4
	PI2 = 4

	fmt.Printf("PI1: %.12f\n", PI1)
	fmt.Printf("PI2: %f\n", PI2)
}
