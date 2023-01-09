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

	word, filenames := ParseArguments()

	var fileList []string
	for _, filename := range filenames {
		fileList = append(fileList, GetMatchingFiles(filename)...)
	}

	if len(fileList) == 0 {
		fmt.Println("No matching files found")
		return
	}

	for _, filename := range fileList {
		fmt.Println(filename)

		wd, _ := os.Getwd()
		filepath := filepath.Join(wd, filename)

		reader, err := ReadFile(filepath)
		if err != nil {
			fmt.Println(err)
			return
		}

		SearchWord(reader, word)
		fmt.Println()
	}
}

// ParseArguments parses command line arguments
func ParseArguments() (word string, filenames []string) {

	args := os.Args[1:]

	if len(args) < 2 {
		fmt.Println("Command line arguments required")
		return
	}

	word = args[0]
	filenames = args[1:]

	return
}

// GetMatchingFiles returns absolute paths matching given file name
func GetMatchingFiles(filename string) []string {

	matches, _ := filepath.Glob(filename)

	return matches
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

	fmt.Println("---------------------------")

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
	fmt.Println("---------------------------")

}
