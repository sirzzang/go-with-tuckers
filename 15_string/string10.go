package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	str1 := "Hello, World!"
	str2 := str1
	fmt.Println(unsafe.Pointer(&str1), &str1)
	fmt.Println(unsafe.Pointer(&str2), &str2)

	// StringHeader의 포인터와 그 길이 값은 복사되었으나, string 자체의 값이 다르다?
	// 16바이트 값만 복사될 뿐, 문자열 데이터는 복사되지 않는다
	strHeader1 := (*reflect.StringHeader)(unsafe.Pointer(&str1))
	strHeader2 := (*reflect.StringHeader)(unsafe.Pointer(&str2))
	fmt.Println(strHeader1)
	fmt.Println(strHeader2)

}
