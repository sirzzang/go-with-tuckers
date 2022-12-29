package main

import "fmt"

func main() {
	array := [5]int{1, 2, 3, 4, 5}
	slice := array[1:3]

	slice = append(slice, 100)
	fmt.Println(array) // [1 2 3 100 5]
}
