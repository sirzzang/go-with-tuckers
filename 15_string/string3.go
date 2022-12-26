package main

import "fmt"

func main() {

	// single character in single quote
	var c rune = '한'

	fmt.Printf("%T\n", c)
	fmt.Printf("%v\n", c) // rune as int32
	fmt.Printf("%c\n", c) // rune as character
	fmt.Printf("'한' == c: %v\n", c == 54620)

}
