package main

import "fmt"

type Stringer interface {
	String() string
}

type Student struct {
	Name string
	Age  int
}

func main() {
	var stringer Stringer
	s := stringer.(*Student) // Student 타입이 Stringer 인터페이스를 구현하지 않음. 컴파일 에러
	fmt.Println(s)
}
