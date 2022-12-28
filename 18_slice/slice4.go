package main

import "fmt"

func main() {
	var slice = []int{1, 2, 3}
	fmt.Println(slice)
	// append(slice, 4) // UnusedExpr
	slice2 := append(slice, 4)
	fmt.Println(slice2)

}
