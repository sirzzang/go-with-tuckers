package main

import "fmt"

func main() {
	str := "Hello 월드"
	runes := []rune(str)

	// iterate bytes
	for i := 0; i < len(str); i++ {
		fmt.Printf("타입: %T, 값: %d, 문자값: %c\n", str[i], str[i], str[i])
	}
	fmt.Println()

	// iterate slice of runes
	for i := 0; i < len(runes); i++ {
		fmt.Printf("타입: %T, 값: %d, 문자값: %c\n", runes[i], runes[i], runes[i])
	}
	fmt.Println()

	// iterate using array
	for _, v := range str {
		fmt.Printf("타입: %T, 값: %d, 문자값: %c\n", v, v, v)
	}
	fmt.Println()
}
