package main

import (
	"fmt"
	"os"
)

type Writer func(string)

func writeHello(writer Writer /* function type */) {
	writer("Hello, world!")
} // func(string) 타입의 함수에 Hello, World!를 인자로 넘김

func main() {
	f, err := os.Create("test2.txt")
	if err != nil {
		fmt.Println("Failed to create test2.txt", err)
		return
	}

	defer f.Close()

	writeHello(func(msg string) {
		fmt.Fprintln(f, msg)
	})

}
