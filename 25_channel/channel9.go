package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {

	wg.Add(1)
	ctx, cancel := context.WithCancel(context.Background())
	go PrintEverySecond(ctx)
	time.Sleep(3 * time.Second)
	cancel() // ctx의 Done이 닫힘

	wg.Wait()
}

func PrintEverySecond(ctx context.Context) {

	tick := time.Tick(time.Second)

	for {
		select {
		case <-ctx.Done():
			wg.Done() // main 고루틴에서 cancel() 호출 시 닫힘
			return
		case <-tick:
			fmt.Println("tick")
		}
	}
}
