package main

import "fmt"

func GetAge() (int, bool) {
	return 28, true
}

func main() {

	// initialization in if statement
	if age, ok := GetAge(); ok && age < 20 {
		fmt.Println("You are young,", age)
	} else if ok && age < 30 {
		fmt.Println("Don't worry,", age)
	} else if ok {
		fmt.Println("Cheer up,", age)
	} else {
		fmt.Println("Something went wrong!")
	}

	// out of initialized scope
	fmt.Println(age)
}
