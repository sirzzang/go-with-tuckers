package main

import (
	"fmt"
	"os"
)

func main() {

	f, err := os.Create("test.txt") // 항상 새로 만듦
	if err != nil {
		fmt.Println("Failed to create test.txt file: ", err)
	}

	defer fmt.Println("Bye, World!")
	defer fmt.Println("반드시 호출됩니다.")
	defer f.Close()
	defer fmt.Println("파일을 닫습니다.")

	fmt.Println("파일에 씁니다.")
	fmt.Fprintln(f, "Hello, World!!!!")
}
