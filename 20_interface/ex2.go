package main

import "fmt"

type Stringer interface {
	String()
}

type Reader interface {
	Read()
}

func CheckAndRun(stringer Stringer) {
	if r, ok := stringer.(Reader); ok {
		r.Read()
	} else {
		fmt.Println("Fail")
	}

}
