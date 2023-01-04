package main

import (
	"fmt"
	"sync"
	"time"
)

type Job interface {
	Do()
}

type SquareJob struct {
	index int
}

func (s *SquareJob) Do() {
	fmt.Printf("%d job start\n", s.index)
	time.Sleep(1 * time.Second)
	fmt.Printf("%d job end: %d \n", s.index, s.index*s.index)
}

var wg sync.WaitGroup

func main() {
	// 작업 배열
	var jobList [10]Job

	// 작업을 jobList 배열에 할당
	for i := 0; i < 10; i++ {
		jobList[i] = &SquareJob{index: i}
	}

	wg.Add(10)
	for i := 10; i > 0; i-- {
		job := jobList[10-i] // (10-i)번째 고루틴은 i번째 job만 할당 받음
		go func() {
			job.Do() // 할당 받은 job을 실행
			wg.Done()
		}()
	}
	wg.Wait()

}
