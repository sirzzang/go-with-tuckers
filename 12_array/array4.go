package main

import "fmt"

func main() {
	var arr [5]float64 = [5]float64{124.0, 124.1, 125.1, 126.1, 122.33}

	for i, v := range arr {
		fmt.Println(i, v, arr[i])
	}

	for _, v := range arr {
		fmt.Println(v)
	}
}
