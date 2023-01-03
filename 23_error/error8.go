package main

type SomeType struct {
	Field string
}

func main() {
	// panic(42) // int
	// panic("unreachable")                       // string type
	// panic(fmt.Errorf("This is error: %d", 42)) // error type
	panic(SomeType{Field: "SomeData"})
}
