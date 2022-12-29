package main

import "fmt"

func main() {
	var slice []int
	idx := 2

	// 반복문
	slice = []int{1, 2, 3, 4, 5, 6}
	fmt.Println(slice)
	slice = append(slice, 0)
	for i := len(slice) - 1; i > idx; i-- {
		slice[i] = slice[i-1]
	}
	slice[idx] = 200
	fmt.Println(slice)

	// append
	slice = []int{1, 2, 3, 4, 5, 6}
	fmt.Println(slice)
	slice = append(slice[:idx], append([]int{200}, slice[idx+1:]...)...)
	fmt.Println(slice)

	// append, copy
	slice = []int{1, 2, 3, 4, 5, 6}
	fmt.Println(slice)
	slice = append(slice, 0)
	copy(slice[idx+1:], slice[idx:])
	slice[idx] = 200
	fmt.Println(slice)

}
