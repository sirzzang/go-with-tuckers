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
	return fmt.Sprintf("Student %s, age %d", s.Name, s.Age)
}

func main() {

	var s Stringer
	fmt.Println(s.String()) // panic
	s = &Student{"sirzzang", 28}
	fmt.Println(s.String())

}
