package main

import "fmt"

func main() {

	str1 := "한글 문자열"
	str2 := str1

	fmt.Printf("&str1: %p, &str2: %p, &str1 == &str2: %v\n", &str1, &str2, &str1 == &str2)
}
