package main

import "fmt"

const FixedCnt int = 3

func main() {
	var VariableCnt int = 5

	a := [FixedCnt]int{1, 2, 3}
	b := [VariableCnt]int{1, 2, 3, 4, 5}
	c := [FixedCnt]int{1, 2, 3, 4, 5}
	d := [5]int{1, 2, 3, 4, 5}

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
}
