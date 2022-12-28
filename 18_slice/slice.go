package main

import "fmt"

func main() {
	var slice []int

	if len(slice) == 0 {
		fmt.Println("empty slice", slice)
	}

	slice[1] = 10 // panic
	fmt.Println(slice)
}
