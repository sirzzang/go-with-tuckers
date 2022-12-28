package main

import "fmt"

func main() {
	var s []int
	var slice = make([]int, 3, 3)
	var slice2 = make([]int, 3, 5)
	var slice3 = []int{1, 2, 3}

	fmt.Println(s, len(s), cap(s))
	fmt.Println(slice, len(slice), cap(slice))
	fmt.Println(slice, len(slice2), cap(slice2))
	fmt.Println(slice, len(slice3), cap(slice3))
}
