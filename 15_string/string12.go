package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	var str string = "Hello"
	strHeader := (*reflect.StringHeader)(unsafe.Pointer(&str))
	fmt.Printf("str Data: %x\n", strHeader.Data)

	str += "World"
	fmt.Printf("str Data: %x\n", strHeader.Data)

	str += "Go"
	fmt.Printf("str Data: %x\n", strHeader.Data)

}
