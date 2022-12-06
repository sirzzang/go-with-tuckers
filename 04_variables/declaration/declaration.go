package main

import (
	"fmt"
	"reflect"
)

func main() {
	
	// basic declaration
	var a int = 1
	fmt.Println(a)

	// declaration without initializaiton
	var b string
	fmt.Println(b)

	// declaration without type
	var c = 3
	fmt.Println(c, reflect.TypeOf(c))

	// short declaration
	d := false
	fmt.Println(d, reflect.TypeOf(d))
}