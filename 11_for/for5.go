package main

import "fmt"

func main() {
	a := 1
	b := 1

OuterFor:
	for ; a <= 9; a++ {
		fmt.Println(a)
		for ; b <= 9; b++ {
			fmt.Println(b)
			if a*b == 45 {
				break OuterFor
			}
		}
		b = 1
		fmt.Println()
	}

	fmt.Printf("%d * %d = %d\n", a, b, a*b)
}
