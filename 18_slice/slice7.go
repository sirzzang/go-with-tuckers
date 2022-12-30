package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {

	slice1 := make([]int, 3, 5)
	fmt.Println(slice1, len(slice1), cap(slice1))

	slice2 := append(slice1, 4, 5)
	fmt.Println(slice2, len(slice2), cap(slice2))
	fmt.Println(slice1, len(slice1), cap(slice1))
	// 헷갈리지 말 것. slice1, slice2가 가리키는 배열 시작 메모리 위치가 Data에 들어가고 그건 일치함
	// slice2가 slice1이 가리키는 배열에서 뒤에 4, 5를 추가했다고 해도, slice1의 len, cap은 그대로임
	// 따라서 이 상태에서 slice1은 뒤에 [0 0 0 4 5]가 아니라, 시작 위치부터 len까지인 [0 0 0]임

	slice1Ptr := (*reflect.SliceHeader)(unsafe.Pointer(&slice1))
	slice2Ptr := (*reflect.SliceHeader)(unsafe.Pointer(&slice2))
	fmt.Println(slice1Ptr.Data, slice1Ptr.Len, slice1Ptr.Cap)
	fmt.Println(slice2Ptr.Data, slice2Ptr.Len, slice2Ptr.Cap)

	slice2[3] = 400
	fmt.Println(slice1) // slice1에는 영향 없음
	fmt.Println(slice2)

	slice2[1] = 200
	fmt.Println(slice1) // slice1도 변함
	fmt.Println(slice2)

	slice1 = append(slice1, -1)
	fmt.Println(slice1)
	fmt.Println(slice1Ptr.Data, slice1Ptr.Len, slice1Ptr.Cap) // 동일한 배열을 가리킴
	fmt.Println(slice2)                                       // slice2도 변함

}
