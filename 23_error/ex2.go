package main

import "fmt"

func sum(nums ...int) int {
	if len(nums) == 0 {
		panic("nums should not be empty")
	}
	result := 0
	for _, v := range nums {
		result += v
	}
	return result
}

func PrintSum() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	fmt.Println(sum(1, 2, 3, 4, 5))
	fmt.Println(sum())
	fmt.Println(sum(1, 2, 3)) // 이건 실행되지 않음
}

func main() {
	PrintSum()
	fmt.Println(("end of main"))
}
