package main

import (
	"fmt"
	"sync"
	"time"
)

type Account struct {
	Balance int
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

func DepositAndWithdraw(account *Account) {
	if account.Balance < 0 {
		panic(fmt.Sprintf("balance must be positive, not %v", account.Balance))
	}

	account.Balance += 1000 // 첫 번째 고루틴이 sleep 중일 때 다른 고루틴이 접근하면 Balance는 증가하지 않을 수 있음
	time.Sleep(time.Millisecond)
	account.Balance -= 1000

}
