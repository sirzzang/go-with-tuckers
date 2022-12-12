package main

import (
	"fmt"
	"unsafe"
)

type Example struct {
	A int8 // 1byte
	B int  // 8bytes
	C int8 // 1byte
	D int  // 8bytes
	E int8 // 1byte
}

type ExampleAligned struct {
	A int
	B int
	C int8
	D int8
	E int8
}

func main() {
	example := Example{1, 2, 3, 4, 5}
	fmt.Println(unsafe.Sizeof(example))

	example2 := ExampleAligned{1, 2, 3, 4, 5}
	fmt.Println(unsafe.Sizeof(example2))

}
