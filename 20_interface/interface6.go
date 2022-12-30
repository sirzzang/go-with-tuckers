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
	if s, ok := stringer.(*Student); ok { // nil 체크까지 해서 type conversion 체크
		fmt.Println(s.String())
	} else {
		fmt.Println("failed")
	}

}
