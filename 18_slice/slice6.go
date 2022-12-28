package main

import "fmt"

func changeArray(array2 [5]int) {
	// 인자로 받은 array 배열의 모든 값을 array2에 복사
	array2[2] = 200
}

func changeSlice(slice2 []int) {
	// 인자로 받은 slice 구조체를 slice2에 복사
	slice2[2] = 200
}

func main() {

	var array = [5]int{1, 2, 3, 4, 5}
	var slice = []int{1, 2, 3, 4, 5}
	fmt.Println(array)
	fmt.Println(slice)

	changeArray(array)
	changeSlice(slice)
	fmt.Println(array)
	fmt.Println(slice)

}
