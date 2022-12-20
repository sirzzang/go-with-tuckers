package main

import "fmt"

type Data struct {
	value int
}

func main() {

	var p1 *Data = &Data{3}
	var p2 *Data = new(Data)
	fmt.Println(p1)
	fmt.Println(p2)
}
