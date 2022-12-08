package main

import "fmt"

func SwitchType(i interface{}) {
	switch v := i.(type) {
	case nil:
		fmt.Println("nil")
	case int:
		fmt.Println("int")
	case bool, string:
		fmt.Println("bool or string")
	default:
		fmt.Println(v)
		fmt.Printf("unknown type %T\n", v)
	}
}

func main() {
	SwitchType("blue")
	SwitchType(23)
	SwitchType(true)
	SwitchType(nil)
	SwitchType([]int{})
	SwitchType([3]int{})
}
