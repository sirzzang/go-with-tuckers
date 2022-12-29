package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {

	array := [6]int{1, 2, 3, 4, 5, 6}
	slice := array[1:3]
	fmt.Println(slice, len(slice), cap(slice)) // 슬라이싱 시 최대 용량은 기존 배열 길이
	fmt.Println()

	// 메모리 공간 비교
	arrayPtr := unsafe.Pointer(&array)
	slicePtr := unsafe.Pointer(&slice)
	fmt.Println(arrayPtr)
	fmt.Println(slicePtr)

	sliceHdr := (*reflect.SliceHeader)(slicePtr)
	fmt.Println(fmt.Sprintf("0x%x", sliceHdr.Data), sliceHdr.Len, sliceHdr.Cap) // 슬라이싱한 위치부터 포인터가 시작함
	fmt.Println()

	// array 요소 변경
	array[0] = 10
	fmt.Println(array)
	fmt.Println(slice, len(slice), cap(slice)) // slice 영향 없음

	array[1] = 20
	fmt.Println(array)
	fmt.Println(slice, len(slice), cap(slice)) // slice도 변경됨
	fmt.Println()

	// slice 요소 추가
	slice = append(slice, 400)
	fmt.Println(array) // array도 변경됨
	fmt.Println(slice, len(slice), cap(slice))
	fmt.Println(fmt.Sprintf("0x%x", sliceHdr.Data), sliceHdr.Len, sliceHdr.Cap) // slice 길이만 변경됨

	slice = append(slice, 500, 600, 700, 800) // slice 용량 부족
	fmt.Println(array)
	fmt.Println(slice, len(slice), cap(slice))
	fmt.Println(fmt.Sprintf("0x%x", sliceHdr.Data), sliceHdr.Len, sliceHdr.Cap) // 새로운 slice
	fmt.Println()

}
