package main

import "fmt"

func main() {

	slice1 := []int{1, 2, 3, 4, 5}

	// slice 생성 후 기존 값 복사하여 복제
	slice2 := make([]int, len(slice1))
	for i, v := range slice1 { // should use copy instead of for loop
		slice2[i] = v
	}
	fmt.Println(slice1, slice2)
	slice2[2] = 300
	fmt.Println(slice1, slice2) // slice1 영향 없음

	// append 함수 이용해서 복제
	slice3 := append([]int{}, slice1...)
	fmt.Println(slice1, slice3)
	slice3[2] = 3000
	fmt.Println(slice1, slice3)

	// copy 함수 이용해서 복제
	slice4 := make([]int, len(slice1))
	slice5 := make([]int, 3, 10)
	slice6 := make([]int, 10)
	cnt4 := copy(slice4, slice1) // 복제된 요소 개수 반환
	cnt5 := copy(slice5, slice1)
	cnt6 := copy(slice6, slice1)
	fmt.Println(cnt4, slice4)
	fmt.Println(cnt5, slice5)
	fmt.Println(cnt6, slice6)

}
