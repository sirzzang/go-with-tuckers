package main

import "fmt"

type account struct {
	balance int
}

func withdrawFunc(a *account, amount int) {
	a.balance -= amount // (*a).balance -= amount
}

func (a *account) withdrawMethod(amount int) {
	a.balance -= amount
}

func main() {
	var a account = account{100}
	fmt.Println(a.balance)

	withdrawFunc(&a, 30)
	fmt.Println(a.balance)

	a.withdrawMethod(30) //	(&a).withdrawMethod(30)
	fmt.Println(a.balance)

	b := &account{200}
	withdrawFunc(b, 40)
	b.withdrawMethod(40)
	fmt.Println(b.balance)
}
