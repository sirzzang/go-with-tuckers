package main

import "fmt"

func enumerate(arr [2][5]int) {
	for _, v := range arr {
		for i, v := range v {
			fmt.Println(i, v)
		}
	}
}

func main() {

	var a = [2][5]int{{1, 2, 3, 4, 5}, {6, 7, 8, 9, 10}}

	var b = [2][5]int{
		{1, 2, 3, 4, 5},
		{6, 7, 8, 9, 10}}

	var c = [2][5]int{
		{1, 2, 3, 4, 5},
		{6, 7, 8, 9, 10},
	}

	var d = [2][5]int{
		{1, 2, 3, 4, 5},
		{6, 7, 8, 9, 10},
	}

	var e = [2][5]int{
		{1, 2, 3, 4, 5},
		{6, 7, 8, 9, 10}}

	enumerate(a)
	fmt.Println()
	enumerate(b)
	fmt.Println()
	enumerate(c)
	fmt.Println()
	enumerate(d)
	fmt.Println()
	enumerate(e)
}
