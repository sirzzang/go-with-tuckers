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
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	go PrintEverySecond(ctx)
	time.Sleep(3 * time.Second)
	// cancel() // 10초 지나기 전에 호출되면 바로 닫힘

	wg.Wait()
}

func PrintEverySecond(ctx context.Context) {

	tick := time.Tick(time.Second)

	for {
		select {
		case <-ctx.Done(): // 컨텍스트에 done 오면 닫힘
			wg.Done() // main 고루틴에서 cancel() 호출 시 닫힘
			return
		case <-tick:
			fmt.Println("tick")
		}
	}
}
