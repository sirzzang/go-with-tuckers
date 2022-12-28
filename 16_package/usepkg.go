package main

import (
	"fmt"
	"go-with-tuckers/16_package/custompkg" // 같은 go 모듈 안의 패키지는 모듈명 아래 위치하도록 해야 함

	"github.com/guptarohit/asciigraph"
	"github.com/tuckersGo/musthaveGo/ch16/expkg"
)

func main() {
	custompkg.PrintCustom()
	expkg.PrintSample()

	data := []float64{3, 4, 5, 6, 9, 7, 5, 8, 5, 10, 2, 7, 2, 5, 6}
	graph := asciigraph.Plot(data)
	fmt.Println(graph)
}
