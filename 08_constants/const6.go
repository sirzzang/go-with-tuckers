package main

import "fmt"

const PI = 3.14
const FloatPi float64 = 3.14

func main() {
	var a int = PI * 100
	fmt.Println(a)

	var a2 string = PI
	fmt.Println(a2)

	var b int = FloatPi * 100
	fmt.Println(b)

}
