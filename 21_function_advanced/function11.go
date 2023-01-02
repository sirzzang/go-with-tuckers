package main

import "fmt"

func main() {
	f := make([]func() int, 3)
	fmt.Println(f)

	for i := 0; i < 3; i++ {
		f[i] = func() int {
			return i
		}
	}

	for i := 0; i < 3; i++ {
		f[i]()
	}
	fmt.Println(f)
}
