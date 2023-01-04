package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	rand.Seed(time.Now().UnixNano())

	wg.Add(2)
	fork := &sync.Mutex{}  // 공유 포크
	spoon := &sync.Mutex{} // 공유 스푼

	go diningProblem("fork first", fork, spoon, "fork", "spoon")
	go diningProblem("spoon first", spoon, fork, "spoon", "fork")
	wg.Wait() // 메인 고루틴 블락

}

func diningProblem(name string, first, second *sync.Mutex, firstName, secondName string) {
	// first, second는 뮤텍스 자원

	for i := 0; i < 100; i++ {
		fmt.Printf("%s 포크와 스푼을 얻으려 한다.\n", name)
		first.Lock() // 첫 번째 획득한 자원 락
		fmt.Printf("%s %s 자원 획득 후 Lock.\n", name, firstName)
		second.Lock() // 두 번째 자원을 획득하려 함
		fmt.Printf("%s %s 자원 획득 후 Lock.\n", name, secondName)

		// 다른 고루틴에서 두 번째 자원을 획득해서 락 걸린 상태라면 여기 아래로 진행할 수 없음
		fmt.Printf("%s 밥을 먹는다.\n", name)
		time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond) // 랜덤한 시간만큼 멈춤

		second.Unlock() // 두 번째 자원 해제
		first.Unlock()  // 첫 번째 자원 해제
	}

	wg.Done()
}
