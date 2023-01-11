package main

import (
	"context"
	"fmt"
	"sync"
)

var wg sync.WaitGroup

type MyString string

func main() {

	wg.Add(1)
	ctx := context.WithValue(context.Background(), MyString("number"), 9)
	go square(ctx)

	wg.Wait()

}

func square(ctx context.Context) {

	v := ctx.Value(MyString("number"))

	if v == nil {
		fmt.Println("no number value")
	} else {
		n := v.(int)
		fmt.Println("square:", n*n)
	}

	wg.Done()

}
