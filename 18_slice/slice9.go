package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	slice1 := []int{1, 2, 3}
	slice1Ptr := (*reflect.SliceHeader)(unsafe.Pointer(&slice1))
	fmt.Println(slice1, unsafe.Pointer(&slice1), slice1Ptr.Data, slice1Ptr.Len, slice1Ptr.Cap)

	slice2 := append(slice1, 4, 5)
	slice2Ptr := (*reflect.SliceHeader)(unsafe.Pointer(&slice2))
	fmt.Println(slice2, unsafe.Pointer(&slice2), slice2Ptr.Data, slice2Ptr.Len, slice2Ptr.Cap) // slice1이 가리키는 배열과 다른 배열 만들어 반환

	slice2 = append(slice2, 6)
	fmt.Println(slice1, slice1Ptr.Data, slice1Ptr.Len, slice1Ptr.Cap) // slice1은 영향 없음
	fmt.Println(slice2, slice2Ptr.Data, slice2Ptr.Len, slice2Ptr.Cap)

	slice1 = append(slice1, 100, 200, 300, 400)                                                // TODO: 길이 늘어나는 방식 뭐지?
	fmt.Println(slice1, unsafe.Pointer(&slice1), slice1Ptr.Data, slice1Ptr.Len, slice1Ptr.Cap) // slice1이 새로운 배열을 가리키게 됨. slice1 식별자가 가리키는 구조체의 주소는 동일
	fmt.Println(slice2, slice2Ptr.Data, slice2Ptr.Len, slice2Ptr.Cap)                          // slice2는 영향 없음
}
