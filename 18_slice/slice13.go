package main

import "fmt"

func main() {
	var slice []int
	idx := 2

	// 반복문
	slice = []int{1, 2, 3, 4, 5, 6}
	for i := idx + 1; i < len(slice); i++ {
		slice[i-1] = slice[i]
	}
	slice = slice[:len(slice)-1]
	fmt.Println(slice)

	// append
	slice = []int{1, 2, 3, 4, 5, 6}
	slice = append(slice[:idx], slice[idx+1:]...)
	fmt.Println(slice)

}
