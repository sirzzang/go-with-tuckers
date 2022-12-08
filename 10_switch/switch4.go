package main

import "fmt"

func GetAge() int {
	return 28
}

func main() {

	switch age := GetAge(); age {
	case 10:
		fmt.Println("Teenager")
	case 30:
		fmt.Println("Thirty")
	default:
		fmt.Println("How old are you?")

	}

	// fmt.Println("Age:", age)

}
