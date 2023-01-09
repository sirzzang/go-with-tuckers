package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

func main() {

	word, filename := ParseArguments()
	filepath := GetFilePath(filename)

	reader, err := ReadFile(filepath)
	if err != nil {
		fmt.Println(err)
		return
	}

	SearchWord(reader, word)

}

// ParseArguments parses command line arguments
func ParseArguments() (word string, filename string) {

	args := os.Args[1:]

	if len(args) < 2 {
		fmt.Println("Command line arguments required")
		return
	} else if len(args) > 2 {
		fmt.Println("Too many command line arguments")
	}

	word = args[0]
	filename = args[1]

	return
}

// GetFilePath returns absolute path of a given file name
func GetFilePath(filename string) string {
	wd, _ := os.Getwd()
	fmt.Println(wd, filename)
	return filepath.Join(wd, filename)
}

// ReadFile reads file from given path
func ReadFile(filename string) (io.Reader, error) {

	data, err := os.Open(filename)
	if err != nil {
		fmt.Printf("Error reading file: %s", err)
		return nil, err
	}

	return data, nil
}

// SearchWord scans text and finds string containing the given word
func SearchWord(reader io.Reader, word string) {

	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanLines)

	cnt := 0
	for scanner.Scan() {

		cnt += 1
		line := scanner.Text()

		if strings.Contains(line, word) {
			fmt.Println(cnt, line)
		}

	}

}
