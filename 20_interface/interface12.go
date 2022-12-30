package main

import "fmt"

type Reader interface {
	Read()
}

type Closer interface {
	Close()
}

type File struct{}

func (f *File) Read() {
	fmt.Println("Read")
}

func ReadFile(reader Reader) {
	if c, ok := reader.(Closer); ok {
		c.Close()
	} else {
		fmt.Println("Converting reader to Closer type failed.")
	}
}

func main() {
	f := &File{}
	ReadFile(f)
}
