package main

import (
	"container/list"
	"fmt"
	"unsafe"
)

type List struct {
	root list.Element
	len  int
}

func main() {

	var list List = List{}
	fmt.Println(list.root.Value, list.root.Next().Value, list.root.Prev().Value)

	fmt.Println(unsafe.Sizeof(list), unsafe.Pointer(&list.root))
	fmt.Println(unsafe.Sizeof(list.root.Value), unsafe.Sizeof(list.root.Next), unsafe.Sizeof(list.root.Prev))
	fmt.Println(unsafe.Sizeof(list.len), unsafe.Pointer(&list.len))
}
