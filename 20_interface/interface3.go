package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

type Reader interface {
	Read()
	Close()
}

type Writer interface {
	Write()
	Close()
}

// ReaderWriter embeds other interfaces
type ReaderWriter interface {
	Reader
	Writer
	Greet()
}

// Reader, Writer, ReaderWriter
type MyStruct struct{}

func (s *MyStruct) Read() {
	fmt.Printf("%s at %p Read\n", reflect.TypeOf(s), unsafe.Pointer(s))
}

func (s *MyStruct) Write() {
	fmt.Printf("%s at %p Write\n", reflect.TypeOf(s), unsafe.Pointer(s))
}

func (s *MyStruct) Close() {
	fmt.Printf("%s at %p Close\n", reflect.TypeOf(s), unsafe.Pointer(s))
}

func (s *MyStruct) Greet() {
	fmt.Println("Should implement this method when it is a ReaderWriter type.")
}

func main() {

	myStruct := &MyStruct{}

	var r Reader = myStruct
	r.Read()
	r.Close()
	// r.Write()
	fmt.Println()

	var w Writer = myStruct
	w.Write()
	w.Close()
	fmt.Println()

	var rw ReaderWriter = myStruct
	rw.Read()
	rw.Write()
	rw.Close()
	fmt.Println()

}
