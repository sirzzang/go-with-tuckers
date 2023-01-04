package main

import "sync"

var mutex sync.Mutex
var wg sync.WaitGroup

type Account struct {
	Balance int
}

func DepositAndWithdraw(account *Account) {
	mutex.Lock()
	defer mutex.Unlock()

	// 자원을 획득해야 아래 로직 실행 가능
	if account.Balance < 0 {
		panic("account.Balance must be positive")
	}
	account.Balance += 1000
	account.Balance -= 1000
	wg.Done()
}

func main() {
	account := &Account{0}
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			DepositAndWithdraw(account)
		}()
	}
	wg.Wait()
}
