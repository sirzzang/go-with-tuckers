package main

import "fmt"

type MyInt int

func (a MyInt) add(b int) int {
	return int(a) + b
}

func main() {
	a := MyInt(3)
	sum := a.add(4)
	fmt.Println(sum)
}
