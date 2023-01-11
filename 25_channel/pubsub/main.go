package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {

	ctx, cancel := context.WithCancel(context.Background())

	wg.Add(4)
	publisher := NewPublisher("Go Publisher", ctx)
	subscriber1 := NewSubscriber("Eraser", ctx)
	subscriber2 := NewSubscriber("Sirzzang", ctx)

	// Publisher 계속 업데이트하면서 구독자 있는지, 발행할 메시지 있는지 확인
	go publisher.Update()

	subscriber1.Subscribe(publisher)
	subscriber2.Subscribe(publisher)

	// Subscriber 계속 업데이트하면서 수신한 메시지 있는지 확인
	go subscriber1.Update()
	go subscriber2.Update()

	// 2초에 한 번씩 메시지 발행
	go func() {
		tick := time.Tick(2 * time.Second)
		for {
			select {
			case <-tick:
				publisher.Publish("Hello, world!")
			case <-ctx.Done():
				wg.Done()
				return
			}
		}
	}()

	// 표준 입력 있으면 종료
	fmt.Scanln()
	cancel()

	wg.Wait()
}
