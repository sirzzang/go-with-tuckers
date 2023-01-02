package main

import "fmt"

func CaptureLoop() {
	f := make([]func(), 3)

	for i := 0; i < 3; i++ {
		fmt.Printf("i: %v at %p\n", i, &i)

		fmt.Println(&i, i)

		// 함수 리터럴 정의
		f[i] = func() { // f[i]의 i는 그 시점에 값으로 바로 평가
			fmt.Println(i)
		}
		fmt.Printf("f[i]: %v at %p\n", f[i], &f[i])
	}

	// 함수 리터럴 호출
	for i := 0; i < 3; i++ {
		f[i]() // 호출 시점에 리터럴 인스턴스 참조 값 가져 옴
	}
}

func main() {
	CaptureLoop()
}
