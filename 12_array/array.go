package main

import "fmt"

func main() {

	var nums [5]int
	fmt.Println("Initialized:", nums)
	nums = [5]int{1, 2, 3, 4, 5}
	fmt.Println("Assigned:", nums)

	days := [3]string{"Sunday", "Monday", "Tuesday"}
	fmt.Println("Initialized and assigned:", days)

	var temps [5]float64 = [5]float64{24.3, 26.7}
	fmt.Println("Initialized and assigned some elements:", temps)

	var arr [5]int = [5]int{2: 2, 3: 4}
	fmt.Println("Initialized and assigned some elements:", arr)

	a := [...]int{1, 2, 3}
	fmt.Println("Array length undefined:", a)

	b := [...]int{1: 1, 9: 2, 10: 4}
	fmt.Println("Array length undefined:", b, "Array length:", len(b))

}
