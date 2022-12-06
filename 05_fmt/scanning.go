package main

import "fmt"

func Scanex1() {
	var a int
	var b int

	// input with scan
	n, err := fmt.Scan(&a, &b)
	if err != nil {
		fmt.Println(n, err)
	} else {
		fmt.Println(n, a, b)
	}

}
