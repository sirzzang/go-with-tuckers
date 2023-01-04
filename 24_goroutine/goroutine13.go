package main

import (
	"fmt"
	"sync"
	"time"
)

var mutex sync.Mutex

type Account struct {
	Balance int
}

func DepositAndWithdraw(account *Account) {
	mutex.Lock() // 뮤텍스 획득

	// 뮤텍스를 획득한 고루틴만 아래 로직에 진입할 수 있음
	defer mutex.Unlock() // 뮤텍스 반환 보장

	if account.Balance < 0 {
		panic(fmt.Sprintf("balance must be positive, not %v", account.Balance))
	}

	account.Balance += 1000
	time.Sleep(time.Millisecond)
	account.Balance -= 1000

}

func main() {
	var wg sync.WaitGroup

	account := &Account{0}

	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			for {
				DepositAndWithdraw(account) // 각 고루틴에서 무한히 호출
			}
		}()
	}
	wg.Wait()

}
