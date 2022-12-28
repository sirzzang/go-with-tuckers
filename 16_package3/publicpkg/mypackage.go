package publicpkg // publicpkg명으로 접근 가능

import "fmt"

const (
	PI = 3.1415   // public
	pi = 3.141592 // private
)

var (
	Screensize   int = 1080
	screenheight int = 960
)

func PublicFunc() {
	const MyConst = 100 // private
	fmt.Println("Public Function:", MyConst)
}

func privateFunc() {
	fmt.Println("Private Function")
}

type MyInt int // public
type myInt int // private

type MyStruct struct {
	Age  int    // public
	name string // private
}

type myStruct struct { // private
	Age  int
	name string
}

func (m myStruct) PrivateMethod() {
	fmt.Println("Private Method")
}
