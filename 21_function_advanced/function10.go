package main

import "fmt"

func CaptureLoop2() {

	f := make([]func(), 3)

	for i := 0; i < 3; i++ {
		v := i // v 변수에 i 값 복사
		fmt.Printf("v: %v at %p, i: %v at %p\n", v, &v, i, &i)

		f[i] = func() {
			fmt.Println(v)
		}
		fmt.Printf("f[i]: %v at %p\n", f[i], &f[i])
	}

	for i := 0; i < 3; i++ {
		f[i]()
	}
}

func main() {
	CaptureLoop2()
}
