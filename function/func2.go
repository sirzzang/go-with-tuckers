package main

import "fmt"

func PrintAvgScore(name string, kor int, math int, eng int) {

	total := kor + math + eng
	avg := float64(total / 3)

	fmt.Printf("%s's Average Score: %.2f\n", name, avg)
}

func main() {

	PrintAvgScore("eraser", 100, 90, 80)
	PrintAvgScore("sirzzang", 99, 91, 92)

}
