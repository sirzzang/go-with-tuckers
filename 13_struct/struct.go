package main

import "fmt"

type House struct {
	Address string
	Size    int
	Price   float64
	Type    string
}

func main() {

	// 초깃값 생략
	var house House
	fmt.Println(house.Address, house.Size, house.Price, house.Type)

	// 모든 필드 초기화
	var house2 House = House{"서울시", 28, 9.80, "아파트"}
	fmt.Println(house2)

	var house3 House = House{
		"서울시",
		28,
		9.80, "아파트"}
	fmt.Println(house3)

	var house4 House = House{
		"서울시", 28, 9.80,
		"아파트", // 닫는 중괄호와 다른 줄에 있으면 쉼표
	}
	fmt.Println(house4)

	house5 := House{
		"서울시",
		28,
		9.80,
		"아파트",
	}
	fmt.Println(house5)

	// 일부 필드 초기화
	var house6 House = House{Size: 28, Price: 9.80}
	fmt.Println(house6)

	var house7 House = House{
		Size: 28, Price: 9.80, // 닫는 중괄호와 다른 줄에 있으면 쉼표
	}
	fmt.Println(house7)

	house8 := House{Size: 28,
		Type: "아파트"}
	fmt.Println(house8)

}
