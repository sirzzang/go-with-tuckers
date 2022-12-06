package main

import "fmt"

func Scanex2() {
	var a int
	var b int

	// input with scanf cause error,
	// when format is invalid
	n, err := fmt.Scanf("%d %d\n", &a, &b)
	if err != nil {
		fmt.Println(n, err)
	} else {
		fmt.Println(n, a, b)
	}

}
