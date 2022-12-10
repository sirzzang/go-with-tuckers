package main

import "fmt"

func iterate() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
	fmt.Println("iterate() done")
}

func iterate2() {
	i := 0
	for ; i < 10; i++ {
		fmt.Println(i)
	}
	fmt.Println("iterate2() done")
}

func iterate3() {
	for i := 0; i < 10; {
		fmt.Println(i)
		i++
	}
	fmt.Println("iterate3() done")
}

func iterate4() {
	i := 0
	for i < 10 {
		fmt.Println(i)
		i++
	}
	fmt.Println("iterate4() done")
}

func iterate5() {
	i := 0
	for i < 10 {
		fmt.Println(i)
		i++
	}
	fmt.Println("iterate5() done")
}

func main() {
	iterate()
	iterate2()
	iterate3()
	iterate4()
	iterate5()
}
