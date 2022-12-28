package main

import "fmt"

func main() {
	var slice = []int{1, 2, 3}

	for i := 0; i < len(slice); i++ {
		fmt.Println(slice[i])
	}

	for i, v := range slice {
		fmt.Printf("%d번째 요소: %d\n", i, v)
	}
}
