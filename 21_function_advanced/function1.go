package main

import "fmt"

func say(num int, msg ...string, num2 float64) {
	for _, s := range msg {
		fmt.Println(s)
	}
}

func main() {
	say(2, "This", "is", "book", 3)
}
