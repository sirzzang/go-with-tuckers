package main

import "fmt"

// named return value
func Divide(a, b int) (result int, success bool) {
	if b == 0 {
		result = 0
		success = false
		return // return statement without values
	}
	result = a / b
	success = true
	return
}

func main() {
	c, success := Divide(11, 3)
	fmt.Println(c, success)
	d, success := Divide(11, 0)
	fmt.Println(d, success)
}
