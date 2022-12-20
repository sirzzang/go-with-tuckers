package main

import "fmt"

type User struct {
	Name string
	Age  int
}

func NewUser(name string, age int) *User {

	var u = User{name, age}
	return &u
}

func main() {

	var p *User = NewUser("sirzzang", 28)
	fmt.Printf("userPointer: %p\n", p)

}
