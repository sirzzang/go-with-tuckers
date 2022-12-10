package main

import "fmt"

func find45(a int) (int, bool) {
	for b := 1; b <= 9; b++ {
		if a*b == 45 {
			return b, true
		}
	}
	return 0, false
}

func main() {
	a := 1
	b := 0

	for ; a <= 9; a++ {
		var flag bool
		if b, flag = find45(a); flag {
			break
		}
	}
	fmt.Printf("%d x %d = %d\n", a, b, a*b)
}
