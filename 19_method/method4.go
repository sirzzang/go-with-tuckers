package main

import "fmt"

type Vertex struct {
	X, Y float64
}

func ScaleFunc(v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func (v Vertex) ScaleMethod(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Vertex{3, 4}
	vPointer := &v
	fmt.Println(v)          // {3, 4}
	vPointer.ScaleMethod(5) // (*vPointer).scaleMethod(5)
	fmt.Println(v)          // {3 4}
}
