package main

import (
	"container/list"
	"fmt"
	"unsafe"
)

func main() {
	v := list.New()
	fmt.Println("v", v, &v, unsafe.Sizeof(v))

	e4 := v.PushBack(4)
	fmt.Println("v", v, &v, unsafe.Sizeof(v))

	e1 := v.PushFront(1)
	fmt.Println("v", v, &v, unsafe.Sizeof(v))

	e3 := v.InsertBefore(3, e4)
	fmt.Println("v", v, &v, unsafe.Sizeof(v))

	e2 := v.InsertAfter(2, e1)
	fmt.Println("v", v, &v, unsafe.Sizeof(v))

	fmt.Println()
	fmt.Println("e1", e1, &e1, unsafe.Sizeof(e1))
	fmt.Println("e2", e2, &e2, unsafe.Sizeof(e2))
	fmt.Println("e3", e3, &e3, unsafe.Sizeof(e3))
	fmt.Println("e4", e4, &e4, unsafe.Sizeof(e4))

	for e := v.Front(); e != nil; e = e.Next() {
		fmt.Print(e.Value, " ")
	}
	fmt.Println()
	for e := v.Back(); e != nil; e = e.Prev() {
		fmt.Print(e.Value, " ")
	}

}
