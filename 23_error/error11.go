package main

import "fmt"

type MyCustomError struct {
	a int
	b int
}

func (e MyCustomError) Error() string {
	return "0으로 나눌 수 없습니다."
}

// f -> g
func f() {

	fmt.Println("f() called")
	g()
	defer fmt.Println("f() ended")

}

// g -> h
func g() {
	defer fmt.Println("g() ended.")
	fmt.Printf("9 / 0 = %d\n", h(9, 0))
}

func h(a, b int) int {
	defer fmt.Println("h() ended.")
	if b == 0 {
		panic(MyCustomError{a: a, b: b})
	}
	return a / b
}

func main() {
	f()
}
