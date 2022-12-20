package main

import "fmt"

type Data struct {
	value int
	arr   [200]int
}

func changeDataPointer(data *Data) {
	data.value = 100
	data.arr[100] = 999
}

func main() {

	var data Data
	fmt.Printf("main 1. data.value: %d\n", data.value)
	fmt.Printf("main 1. data.arr[100]: %d\n", data.arr[100])

	changeDataPointer(&data)
	fmt.Printf("main 2. data.value: %d\n", data.value)
	fmt.Printf("main 2. data.arr[100]: %d\n", data.arr[100])

}
