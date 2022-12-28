package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {

	// nil slice
	var slice []int
	slicePtr := unsafe.Pointer(&slice)
	sliceHeader := (*reflect.SliceHeader)(slicePtr)
	fmt.Println(slicePtr, sliceHeader.Data)
	fmt.Println(slice == nil, slice, len(slice), cap(slice))
	if len(slice) == 0 {
		fmt.Println("empty slice", slice)
	}
	// fmt.Println(slice[1])
	// slice[1] = 10 // panic when not initialized
	fmt.Println("iteration start")
	for i, v := range slice {
		fmt.Println(i, v)
	}
	fmt.Println("iteration done")
	fmt.Println(slice[:], slice[:] == nil)
	slice = append(slice, 1)
	fmt.Println(append(slice, 1))
	fmt.Println(slicePtr, sliceHeader.Data) // // new underlying array

	// empty slice
	fmt.Println("===== initialize =====")
	slice = []int{}
	fmt.Println(slicePtr, sliceHeader.Data)
	fmt.Println(slice == nil, slice, len(slice), cap(slice))
	if len(slice) == 0 {
		fmt.Println("empty slice", slice)
	}
	slice[1] = 10 // panic when initialized but length is 0

	slice1 := make([]int, 0)
	slice1Ptr := unsafe.Pointer(&slice1)
	slice1Header := (*reflect.SliceHeader)(slice1Ptr)
	fmt.Println(slice1Ptr, slice1Header.Data)
	fmt.Println(slice1 == nil, slice1, len(slice1), cap(slice1))

	fmt.Println("=====")
	slice2 := make([]int, 0)
	slice2Ptr := unsafe.Pointer(&slice2)
	slice2Header := (*reflect.SliceHeader)(slice2Ptr)
	fmt.Println(slice2Ptr, slice2Header.Data)
	fmt.Println(slice2 == nil, slice2, len(slice2), cap(slice2))

}
