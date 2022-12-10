package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	stdin := bufio.NewReader(os.Stdin) // standard input

	for {
		fmt.Println("Input: ")

		var number int
		_, err := fmt.Scanln(&number) // one line input

		if err != nil {
			fmt.Println("Input type expected to be an integer.")
			stdin.ReadString('\n') // empty buffer
			continue
		}

		fmt.Printf("Your input value is %d\n", number)

		if number%2 == 0 {
			fmt.Println("Even number. Bye.")
			break
		}
		fmt.Println("Odd number. Input again.")
	}
}
