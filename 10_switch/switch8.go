package main

import "fmt"

func main() {

	a := 3

	switch a {
	case 1:
		fmt.Println("1")
	case 2:
		fmt.Println("2")
	case 3:
		fmt.Println("3")
		fallthrough
	case 4: // 실행
		fmt.Println("4")
		fallthrough
	case 5: // 여기까지 실행
		fmt.Println("5")
	}
}
