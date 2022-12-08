package main

import "fmt"

func main() {
	temp := 18

	// switch as if
	switch {
	case temp < 10, temp > 30:
		fmt.Println("바깥 활동하기에 좋은 날씨가 아닙니다.")
	case temp >= 15 && temp < 25: // here
		fmt.Println("야외 활동하기 좋은 날씨입니다.")
	case temp >= 10 && temp < 20: // not evaluated
		fmt.Println("약간 추울 수 있으니 가벼운 겉옷을 준비하세요.")
	default:
		fmt.Println("따뜻합니다.")
	}
}
