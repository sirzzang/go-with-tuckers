package main

import (
	"fmt"
	"sync"
	"time"
)

type Car struct {
	Body  string
	Tire  string
	Color string
}

var wg sync.WaitGroup
var startTime = time.Now()

func main() {

	tireCh := make(chan *Car)
	paintCh := make(chan *Car)

	// 차를 만들기 위해서는 3개의 일이 끝나야 한다.
	fmt.Println("Start car factory")
	wg.Add(3)

	// body 담당 고루틴은 MakeBody 일만 한다.
	go MakeBody(tireCh)

	// tire 담당 고루틴은 InstallTire 일만 한다. MakeBody 일이 끝나야 된다.
	// tireCh 채널 타입 변수로 진행 상황을 관리한다.
	go InstallTire(tireCh, paintCh)

	// paint 담당 고루틴은 PaintCar 일만 한다. MakeBody, InstallTire가 끝나야 된다.
	// tireCh, paintCh 채널 타입 변수로 진행 상황을 관리한다.
	go PaintCar(paintCh)

	wg.Wait()
	fmt.Println("Close car factory")

}

// tireCh 생산자
func MakeBody(tireCh chan *Car) {

	tick := time.Tick(time.Second)
	after := time.After(10 * time.Second)
	for {
		select {
		case <-tick: // 1초에 한 번씩 차를 만듦
			car := &Car{}
			car.Body = "Sports car" // Body를 만든다
			tireCh <- car           // 만든 차를 channel에 넣는다 -> install 작업에서 사용된다
		case <-after:
			close(tireCh) // 채널을 닫는다
			wg.Done()     // 작업 완료를 알린다
			return
		}
	}
}

// tireCh 소비자, paintCh 생산자
func InstallTire(tireCh chan *Car, paintCh chan *Car) {

	for car := range tireCh {
		time.Sleep(time.Second)
		car.Tire = "Sports tire" // Tire를 설치한다
		paintCh <- car           // 만든 차를 channel에 넣는다 -> paint 작업에서 사용된다
	}
	wg.Done()      // 작업 완료를 알린다
	close(paintCh) // 채널을 닫는다
}

func PaintCar(paintCh chan *Car) {
	for car := range paintCh {
		time.Sleep(time.Second)
		car.Color = "Sports color"                                                                     // Paint를 칠한다
		duration := time.Now().Sub(startTime)                                                          // 경과 시간
		fmt.Printf("%.2f complete car: %s %s %s\n", duration.Seconds(), car.Body, car.Tire, car.Color) // 고루틴별로 작업시간 측정
	}

	wg.Done() // 작업 완료를 알린다
}
