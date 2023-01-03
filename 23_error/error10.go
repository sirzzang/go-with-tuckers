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
	defer func() {
		if r, ok := recover().(MyCustomError); ok {
			fmt.Printf("recover panic:%v, %v\n", r.Error(), r)
		}
	}()

	g()
	fmt.Println("f() ended")

}

// g -> h
func g() {
	fmt.Printf("9 / 3 = %d\n", h(9, 3))
	fmt.Printf("9 / 0 = %d\n", h(9, 0))
}

func h(a, b int) int {
	if b == 0 {
		panic(MyCustomError{a: a, b: b})
	}
	return a / b
}

func main() {
	f()
	fmt.Println("프로그램 실행")
	f()
	fmt.Println("running program")
}
