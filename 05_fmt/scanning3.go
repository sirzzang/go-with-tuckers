package main

import "fmt"

func Scanex3() {
	var a int
	var b int

	// input with scanln should be exited,
	// with enter key pressed
	n, err := fmt.Scanln(&a, &b)
	if err != nil {
		fmt.Println(n, err)
	} else {
		fmt.Println(n, a, b)
	}

}
