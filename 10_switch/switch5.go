package main

import "fmt"

func GetAge() int {
	return 28
}

func main() {

	switch age := GetAge(); true {
	case age >= 10 && age < 20:
		fmt.Println("Teenager")
	case age < 30:
		fmt.Println("Beautiful Age")
	case age < 50:
		fmt.Println("Nice Age")
	default:
		fmt.Println("How old are you?")

	}

	// fmt.Println("Age:", age)

}
