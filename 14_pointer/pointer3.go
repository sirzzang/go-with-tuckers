package main

import "fmt"

type Data struct {
	value int
	arr   [200]int
}

func changeData(data Data) {
	fmt.Printf("changeData 1. data.value: %d\n", data.value)
	fmt.Printf("changeData 1. data.arr[100]: %d\n", data.arr[100])

	data.value = 100
	data.arr[100] = 999

	fmt.Printf("changeData 2. data.value: %d\n", data.value)
	fmt.Printf("changeData 2. data.arr[100]: %d\n", data.arr[100])
}

func main() {

	var data Data
	fmt.Println(&data, &data == nil)
	fmt.Printf("main 1. data.value: %d\n", data.value)
	fmt.Printf("main 1. data.arr[100]: %d\n", data.arr[100])

	changeData(data)
	fmt.Printf("main 2. data.value: %d\n", data.value)
	fmt.Printf("main 2. data.arr[100]: %d\n", data.arr[100])

}
