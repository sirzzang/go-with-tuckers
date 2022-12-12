package main

import (
	"fmt"
	"unsafe"
)

type Padding struct {
	A int8    // 1, 8
	B int     // 8, 16
	C float64 // 8, 24
	D uint16  // 2, 32
	E int     // 8, 40
	F float32 // 4,
	G int8    // 1, 48
}

type Padding2 struct {
	B int     // 8, 8
	C float64 // 8, 16
	E int     // 8, 24
	F float32 // 4
	D uint16  // 2
	A int8    // 1
	G int8    // 1, 32
}

func main() {
	example := Padding{1, 2, 3.0, 4, 5, 6.0, 7}
	fmt.Println(unsafe.Sizeof(example)) // 32

	example2 := Padding2{0.0, 1, 2, 3, 4, 5.0, 6}
	fmt.Println(unsafe.Sizeof(example2)) // 32

}
