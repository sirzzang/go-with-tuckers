package main

import (
	"bufio"
	"errors"
	"fmt"
	"strconv"
	"strings"
)

func multipleFromString(eq string) (int, error) { // 래핑한 에러는 scanner error 혹은 strconv 에러임

	// read first int value
	scanner := bufio.NewScanner(strings.NewReader(eq))
	scanner.Split(bufio.ScanWords) // scan by word

	// initialize pointer
	pos := 0

	a, n, err := readNextInt(scanner)
	if err != nil {
		return 0, fmt.Errorf(
			"Failed to readNextInt(), pos: %d, err: %w", pos, err) // wrap error
	}

	// move pointer to next scan position
	pos += n + 1

	// read second int value
	b, n, err := readNextInt(scanner)
	if err != nil {
		return 0, fmt.Errorf(
			"Failed to readNextInt(), pos: %d, err: %w", pos, err) // wrap error
	}

	return a * b, nil

}

func readNextInt(scanner *bufio.Scanner) (int, int, error) {

	if !scanner.Scan() { // when false, raise custom error
		return 0, 0, fmt.Errorf("Failed to scan")
	}

	word := scanner.Text()
	number, err := strconv.Atoi(word) // NumError
	if err != nil {
		return 0, 0, fmt.Errorf(
			"Failed to convert word to int, word: %s, err: %w", word, err) // wrap error
	}
	return number, len(word), nil
}

func readEq(eq string) {
	// returns multiple of two numbers read from given string
	result, err := multipleFromString(eq)
	if err == nil {
		fmt.Println(result)
	} else {
		fmt.Println(err)
		var numError *strconv.NumError // unwrap이 포인터리시버 사용
		if errors.As(err, &numError) {
			fmt.Println("NumberError Raised: ", numError)
		}
	}
}

func main() {
	readEq("123 123")
	fmt.Println()
	readEq("123 abc")
	fmt.Println()
	readEq("abc 123")
	fmt.Println()
	readEq("")
	fmt.Println()
	readEq("123 123 123")
	fmt.Println()
	readEq(" 123 123")
}
