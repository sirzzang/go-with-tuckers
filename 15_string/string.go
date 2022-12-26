package main

import "fmt"

func main() {

	// double quote vs. back quote
	str1 := "Hello\t'World'\n"
	str2 := `Hello\t'World'\n`

	// escape character required in double quotes
	str3 := "Hello\\t'World'\\n"
	str4 := "Hello\\t\"World\""

	fmt.Println(str1)
	fmt.Println(str2)
	fmt.Println(str3)
	fmt.Println(str4)

}
