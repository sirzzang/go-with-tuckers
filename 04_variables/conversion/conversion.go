package main

import "fmt"

func main() {

	a := 3
	var b float64 = 3.65
	fmt.Println(b)

	// fmt.Println(a + b) // invalid operation: a + b (mismatched types int and float64)

	// convert float into int
	var c int = int(b)
	fmt.Println(c)
	fmt.Println(a + c)

	// convert int into float64
	d := float64(a * c)
	fmt.Println(d)

	// convert float64 into int64
	var e int64 = 7
	f := int64(d) * e
	fmt.Println(f)

	// convert float64 into int, 
	// value may vary depending on conversion order
	var g int = int(b * 3)
	var h int = int(b) * 3
	fmt.Println(g)
	fmt.Println(h)

	// convert int16 into int8,
	// 1 byte is lost
	var i int16 = 3456
	var j int8 = int8(i)
	fmt.Println(i) // 3456
	fmt.Println(j) // -128

	// decimal points are missed
	var f1 float32 = 1234.523
	var f2 float32 = 3456.123
	var f3 float32 = f1 * f2 * 3
	var f4 float32 = f1 * f2
	var f5 float64 = float64(f1 * f2)
	var f6 float64 = float64(f1) * float64(f2)
	fmt.Println(f3) // 1.2799989e+07(exact value: 12799990.002987) 
	fmt.Println(f4) // 4.266663e+06
	fmt.Println(f5) // 4.266663e+06
	fmt.Println(f6) // 4.266663216691017e+06

}