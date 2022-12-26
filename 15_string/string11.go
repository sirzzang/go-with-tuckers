package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	var str string = "Hello World"
	fmt.Println(str)

	str = "How are you?"
	fmt.Println(str)

	// str[2] = 'a'
	// fmt.Println(str)

	var slice []byte = []byte(str)
	fmt.Println(slice)

	strHeader := (*reflect.StringHeader)(unsafe.Pointer(&str))
	sliceHeader := (*reflect.SliceHeader)(unsafe.Pointer(&slice))
	fmt.Println(strHeader)
	fmt.Println(sliceHeader)

	slice[2] = 'a'
	fmt.Println(slice)

	var str2 string = string(slice)
	fmt.Println(str2)
}
