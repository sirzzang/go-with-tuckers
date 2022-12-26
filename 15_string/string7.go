package main

import "fmt"

func main() {

	str1 := "Hello"
	str2 := "World"
	str3 := "Hell"
	str4 := "Hello"

	// add strings
	str5 := str1 + ", " + str2
	fmt.Println(str5)

	// compare strings
	fmt.Printf("%s == %s : %v\n", str1, str3, str1 == str3)
	fmt.Printf("%s == %s : %v\n", str1, str4, str1 == str4)
	fmt.Printf("&str1: %p, &str4: %p, &str1 == &str4: %v\n", &str1, &str4, &str1 == &str4)

}
