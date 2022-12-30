package main

import "fmt"

func PrintVal(v interface{}) { // 모든 타입의 변수를 인자로 받음
	switch t := v.(type) {
	case int:
		fmt.Printf("v is int, %d\n", int(t))
	case float64:
		fmt.Printf("v is float64, %f\n", float64(t))
	case string:
		fmt.Printf("v is string, %v\n", string(t))
	default:
		fmt.Printf("v is unknown type %T, value %v\n", t, t)
	}
}

type Person struct {
	Name string
	Age  int
}

func main() {
	PrintVal(10)
	PrintVal(3.141592)
	PrintVal("Hello World")
	PrintVal(Person{"sirzzang", 28})
}
