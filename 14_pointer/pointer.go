package main

import (
	"fmt"
)

func main() {

	var a int = 500

	// int형 포인터 변수 p에 a의 메모리 주소를 값으로 대입
	var p *int
	p = &a
	fmt.Printf("p의 값: %p\n", p)          // 0x1400012c008
	fmt.Printf("p의 값이 가리키는 값: %d\n", *p) // 500

	// int형 포인터 변수 p가 가리키는 메모리 공간의 값 변경
	*p = 200
	fmt.Printf("p의 값: %p\n", p)          // 0x1400012c008
	fmt.Printf("p의 값이 가리키는 값: %d\n", *p) // 200
	fmt.Printf("&p의 값: %p\n", &p)        // 0x1400011c018

	var p2 **int
	p2 = &p
	fmt.Printf("&&p의 값: %p\n", &p2) // 0x1400011c0

}
