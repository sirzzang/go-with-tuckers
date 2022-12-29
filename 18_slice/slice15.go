package main

import (
	"fmt"
	"sort"
)

func main() {
	s := []int{5, 1, 6, 2, 9, 3, 22}
	fmt.Println(s)
	sort.Ints(s) // python sort
	fmt.Println(s)
}
