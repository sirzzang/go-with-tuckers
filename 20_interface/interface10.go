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

func ReadFile(reader Reader) { // File 타입 인스턴스를 Reader 인터페이스 타입으로 사용함. 컴파일 에러 발생하지 않음
	c := reader.(Closer) // Reader 인터페이스 타입을 Closer 인터페이스 타입으로 변환하고자 함. 컴파일 에러 발생하지 않음
	c.Close()            // 런타임 에러 발생
}

func main() {
	f := &File{}
	ReadFile(f)
}
