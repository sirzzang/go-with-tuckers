package main

import "fmt"

const a int = iota
const b int = iota

const (
	Red   int = iota
	Blue  int = iota
	Green int = iota
)

const (
	C1 int = iota + 1
	C2 int = iota
	C3 int = iota
)

const (
	Bit1 int = 3 << iota // iota가 0이므로 의미 없음
	Bit2 int = 3 << iota // 00000110
	Bit3 int = 3 << iota // 00001100
)

const (
	Bit11 int = iota << 3 // iota가 0이므로 의미 없음
	Bit22 int = iota << 3 // 00001000
	Bit33 int = iota << 3 // 00010000
)

const (
	BitFlag1 uint = 1 << iota // 00000001
	BitFlag2                  // 00000010
	BitFlag3                  // 00000100
)

const (
	BitFlag11 int = iota << 1 // 00000000
	BitFlag22                 // 00000010
	BitFlag33                 // 00001000
	BitFlag44                 // 00000110
)

func main() {
	fmt.Println(a, b)
	fmt.Println(Red, Blue, Green)
	fmt.Println(C1, C2, C3)
	fmt.Println(Bit1, Bit2, Bit3)
	fmt.Println(Bit11, Bit22, Bit33)
	fmt.Println(BitFlag1, BitFlag2, BitFlag3)
	fmt.Println(BitFlag11, BitFlag22, BitFlag33, BitFlag44)
}
