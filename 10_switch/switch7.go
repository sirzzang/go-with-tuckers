package main

import "fmt"

func main() {

	a := 3

	switch a {
	case 1:
		fmt.Println("1")
		break
	case 2:
		fmt.Println("2")
		break // does not break
	case 3:
		fmt.Println("3")
	case 4:
		fmt.Println("4")
		break // already break
	case 5:
		fmt.Println("5")
	}
}
