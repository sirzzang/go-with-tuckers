package main

import (
	"fmt"
	"reflect"
)

func main() {
	const Gravity = 9.80665
	fmt.Println(reflect.TypeOf(Gravity))

	var b = Gravity * 1000000
	fmt.Println(b, reflect.TypeOf(b))
}
