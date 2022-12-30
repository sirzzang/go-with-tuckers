package main

import "fmt"

type Stringer interface {
	String() string
}

type Student struct {
	Name string
	Age  int
}

func (s *Student) String() string {
	return fmt.Sprintf("Student(%s, %d)", s.Name, s.Age)
}

func main() {
	var stringer Stringer
	s := stringer.(*Student) // Student 타입이 Stringer 인터페이스를 구현함. 그런데 nil이라 런타임 에러
	fmt.Println(s.String())

}
