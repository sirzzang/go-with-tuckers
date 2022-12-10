package main

import "fmt"

func main() {

	a := 1
	b := 1
	flag := false

	for ; a <= 9; a++ {
		for ; b <= 9; b++ {
			if a*b == 45 { // target case found
				flag = true
				break
			}
		}

		if flag { // break outer loop
			break
		}

		b = 1 // set back to original value when not found

	}

	fmt.Printf("%d x %d = %d\n", a, b, a*b)

}
