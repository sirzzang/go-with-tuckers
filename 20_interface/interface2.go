package main

import (
	"fmt"
	"unsafe"
)

type Stringer interface {
	String() string
}

type Student struct {
	Name string
}

// Student implements Stringer interface
func (s *Student) String() string {
	return fmt.Sprintf("Student %s", s.Name)
}

type User struct {
	Name string
	Age  int
}

// User implements Stringer interface
func (u User) String() string {
	return fmt.Sprintf("User %s", u.Name)
}

func main() {
	var stringer1 Stringer
	fmt.Printf("stringer1 type: %T size: %d address: %p\n", stringer1, unsafe.Sizeof(stringer1), unsafe.Pointer(&stringer1))

	s := &Student{"sirzzang"}
	stringer1 = s
	fmt.Printf("s type: %T size: %d address: %p\n", s, unsafe.Sizeof(s), unsafe.Pointer(s))
	fmt.Printf("stringer 1 type: %T size: %d address: %p\n", stringer1, unsafe.Sizeof(stringer1), unsafe.Pointer(&stringer1))
	fmt.Println()

	var stringer2 Stringer
	fmt.Printf("stringer2 type: %T size: %d address: %p\n", stringer2, unsafe.Sizeof(stringer2), unsafe.Pointer(&stringer2))

	u := User{"eraser", 28}
	stringer2 = u
	fmt.Printf("u type: %T size: %d address: %p\n", u, unsafe.Sizeof(u), unsafe.Pointer(&u))
	fmt.Printf("stringer2 type: %T size: %d address: %p\n", stringer2, unsafe.Sizeof(stringer2), unsafe.Pointer(&stringer2))

}
