package main

import "fmt"

func Print(args ...interface{}) string {
	var s string
	for _, arg := range args {
		switch arg.(type) {
		case bool:
			val := arg.(bool)
			s = fmt.Sprintf("bool %v", val)
			fmt.Println(s)
		case float64:
			val := arg.(float64)
			s = fmt.Sprintf("float64 %v", val)
			fmt.Println(s)
		case int:
			val := arg.(int)
			s = fmt.Sprintf("int %v", val)
			fmt.Println(s)
		default:
			s = fmt.Sprintf("unknown %v", arg)
			fmt.Println(s)
		}
	}
	return s
}

func main() {
	Print(1, 2, "hello", 3.141592, "world")
}
