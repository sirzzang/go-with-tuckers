package main

import "fmt"

func Printex() {

	var a int = 10
	var b int = 20
	var c int = 123456789
	var f float64 = 1234.123456789

	// print without new line
	fmt.Print("a: ", a, "b: ", b, "f: ", f)

	// print with new line
	fmt.Println("a: ", a, "b: ", b, "f: ", f)

	// print with formatting
	fmt.Printf("a: %d, b: %d, f: %f", a, b, f)

	// print int values with designated width
	fmt.Printf("%5d, %5d\n", a, b)
	fmt.Printf("%05d, %05d\n", a, b)
	fmt.Printf("%-5d, %-05d\n", a, b)
	fmt.Printf("%5d, %5d\n", c, c)
	fmt.Printf("%05d, %05d\n", c, c)
	fmt.Printf("%-5d, %-05d\n", c, c)

	// print float values with designated points
	fmt.Printf("%08.2f\n", f) // 01234.12
	fmt.Printf("%08.2g\n", f)
	fmt.Printf("%8.5g\n", f)
	fmt.Printf("%f\n", f) // default %.6f

	// print string with special characters
	str := "Hello\tGo\t\tWorld\n\"Go\"is Awesome!\n\n"
	fmt.Print(str)
	fmt.Printf("%s", str)
	fmt.Printf("%q", str)
}
