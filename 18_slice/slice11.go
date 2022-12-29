package main

import "fmt"

func PrintSliceInfo(slice []int) {
	fmt.Println(slice, len(slice), cap(slice))
}

func main() {

	slice1 := []int{1, 2, 3, 4, 5, 6}
	slice2 := slice1[1:3]
	PrintSliceInfo(slice1)
	PrintSliceInfo(slice2)
	fmt.Println()

	slice3 := slice1[0:3]
	slice4 := slice3[:3] // 시작 인덱스 생략
	PrintSliceInfo(slice3)
	PrintSliceInfo(slice4)
	fmt.Println()

	// slice5 := slice1[3:len(slice1)] // 포맷팅에 의해 unneeded index 생략
	slice6 := slice1[3:] // 끝 인덱스 생략
	PrintSliceInfo(slice6)
	fmt.Println()

	slice7 := slice1[:] // 전체 슬라이싱
	PrintSliceInfo(slice7)
	fmt.Println()

	// 슬라이스 용량 지정
	slice8 := slice1[1:3:3]
	// slice9 := slice1[1:5:3] // SwappedSliceIndices 에러
	slice10 := slice1[1:3:4]
	PrintSliceInfo(slice8)
	PrintSliceInfo(slice10)
	fmt.Println()

}
