package main

import "fmt"

// f -> g
func f() {

	fmt.Println("f() called")
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("recover panic:", r)
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
		panic("0으로 나눌 수 없습니다.")
	}
	return a / b
}

func main() {
	f()
	fmt.Println("프로그램 실행")
	f()
	fmt.Println("running program")
}
