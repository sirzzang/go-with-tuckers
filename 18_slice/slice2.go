package main

import (
	"fmt"
	"reflect"
)

func main() {

	// slice 초기화
	var slice1 []int = []int{1, 2, 3}
	var slice2 = []int{1, 2, 3, 4}
	slice3 := []int{1, 5: 2, 10: 3}
	var slice4 = make([]int, 3)
	fmt.Println(slice1)
	fmt.Println(slice2)
	fmt.Println(slice3) // 요소별 위치 정해서 초기화
	fmt.Println(slice4) // default value로 초기화

	// slice vs. array
	var array = [...]int{1, 2, 3}
	var slice = []int{1, 2, 3}
	fmt.Println(reflect.TypeOf(array)) // [3]int: 배열
	fmt.Println(reflect.TypeOf(slice)) // []int: 슬라이스
}
