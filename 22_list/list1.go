package main

import (
	"container/list"
	"fmt"
	"unsafe"
)

func main() {
	var list *list.List = list.New()

	root := list.PushFront(1)
	tail := list.PushBack(5)
	mid := list.InsertAfter(3, root)
	list.InsertBefore(4, tail)
	list.InsertBefore(2, mid)

	for e := list.Front(); e != nil; e = e.Next() {
		fmt.Println(unsafe.Sizeof(e), e.Value)
	}
	fmt.Println()
	fmt.Println(list.Len())
	fmt.Println(list.Front(), root, list.Front() == root)
	fmt.Println("root", root, unsafe.Sizeof(root), unsafe.Sizeof(list.Front()))
	fmt.Println("root.Value", root.Value, unsafe.Sizeof(root.Value))
	// fmt.Println("root.Prev", root.Prev().Value, unsafe.Sizeof(root.Prev().Value))
	fmt.Println("root.Prev", root.Prev(), unsafe.Sizeof(root.Prev()))
	fmt.Println("root.Next", root.Next().Value, unsafe.Sizeof(root.Next().Value))

}
