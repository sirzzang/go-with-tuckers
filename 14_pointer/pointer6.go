package main

import "fmt"

type Data struct {
	value int
}

func main() {

	var p1 *Data = &Data{}
	var p2 *Data = p1
	var p3 *Data = p1
	fmt.Println(&p1)
	fmt.Println(&p2)
	fmt.Println(&p3)
	fmt.Printf("p1 == p2: %v\n", p1 == p2)
	fmt.Printf("p2 == p3: %v\n", p2 == p3)
	fmt.Printf("p3 == p1: %v\n", p3 == p1)

	var p4 Data
	var p5 Data = p4
	var p6 Data = p4
	fmt.Println(&p4)
	fmt.Println(&p5)
	fmt.Println(&p6)
	fmt.Printf("p4 == p5: %v\n", p4 == p5)
	fmt.Printf("p5 == p6: %v\n", p5 == p6)
	fmt.Printf("p6 == p4: %v\n", p6 == p4)
}
