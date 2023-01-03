package main

import (
	"bufio"
	"fmt"
	"os"
)

func ReadFile(filename string) (string, error) {

	// open file
	file, err := os.Open(filename)
	if err != nil {
		fmt.Printf("readfile %v\n", err)
		return "", err // 에러 발생 시 에러 반환
	}
	defer file.Close()

	// read from file
	reader := bufio.NewReader(file)
	line, _ := reader.ReadString('\n') // delimiter 만났을 때이므로 에러 무시
	return line, nil

}

func WriteFile(filename string, line string) error {
	// create file
	file, err := os.Create(filename)
	if err != nil {
		return err // 에러 발생 시 에러 반환
	}
	defer file.Close()

	// write to file
	_, err = fmt.Fprintln(file, line) // 쓴 바이트 반환값 무시
	return err
}

const filename = "data.txt"

func main() {

	// try to read file
	line, err := ReadFile(filename)

	// when error reading file, try to write file
	if err != nil {
		fmt.Printf("%v\n", err)
		err = WriteFile(filename, "No data.txt file, so write it.")

		// when error writing file, return
		if err != nil {
			fmt.Println(err) // 에러 출력
			fmt.Printf("%v\n", err)
			return
		}

		// try to read file again
		line, err = ReadFile(filename)

		// when error reading file, return
		if err != nil {
			fmt.Println(err)
			return
		}

	}

	// print read line
	fmt.Println(line)

}
