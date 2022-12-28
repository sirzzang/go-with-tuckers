package exinit

import "fmt"

// initialized when package imported
var (
	a = c + b // 3. a = 9
	b = f()   // 1. d = 4, b = 4
	c = f()   // 2. d = 5, c = 5
	d = 3
)

// called when package imported
func init() {
	d++ // 4. d = 6
	fmt.Println("init() called:", d)
}

func f() int {
	d++
	fmt.Println("f() called:", d)
	return d
}

func PrintD() {
	fmt.Printf("a: %d, b: %d, c: %d, d: %d\n", a, b, c, d)
	fmt.Println("d:", d)
}
