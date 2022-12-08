package main

import "fmt"

func main() {
	day := "thursday"

	switch day {
	case "monday", "tuesday": // compare multiple values
		fmt.Println("주초입니다.")
	case "wednesday", "thursday":
		fmt.Println("주중입니다.")
	case "friday":
		fmt.Println("주말이 얼마 남지 않았습니다.")
	}
}
