package main

import "fmt"

func main() {

	i := 0
	fmt.Println(1, &i, i)

	// 함수 리터럴이 정의되는 시점
	f := func() {
		fmt.Println("in f", &i)
		i += 10
	}

	i++

	f() // 함수 리터럴이 호출되는 시점에

	fmt.Println(i, &i)
}
