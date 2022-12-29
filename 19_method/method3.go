package main

import (
	"fmt"
	"unsafe"
)

type account struct {
	balance   int    // 8
	firstName string // 16
	lastName  string // 16
}

func (a *account) withdrawPointer(amount int) {
	fmt.Println("withdrawPointer a memory: ", unsafe.Pointer(a))
	a.balance -= amount
}

func (a account) withdrawValue(amount int) {
	fmt.Println("withdrawValue a memory: ", unsafe.Pointer(&a))
	a.balance -= amount
}

func (a account) withdrawValueAndReturn(amount int) account {
	fmt.Println("withdrawValueAndReturn a memory: ", unsafe.Pointer(&a))
	a.balance -= amount
	return a
}

func main() {

	accountPtr := &account{1000, "sir", "zzang"}
	fmt.Println(unsafe.Pointer(accountPtr))
	fmt.Println(accountPtr.balance)

	// 포인터 변수 포인터 타입 메서드 호출
	accountPtr.withdrawPointer(100)
	fmt.Println(accountPtr.balance)

	// 포인터 변수 값 타입 메서드 호출
	accountPtr.withdrawValue(100) // (*accountPtr).withdrawValue(100)
	fmt.Println(accountPtr.balance)

	// 포인터 변수 값 타입 메서드 호출 후 값 타입 반환
	accountVal := accountPtr.withdrawValueAndReturn(100)
	fmt.Println(unsafe.Pointer(&accountVal))
	fmt.Println(accountVal.balance)
	fmt.Println(accountPtr.balance)

	// 값 변수 포인터 타입 메서드 호출
	accountVal.withdrawPointer(100) // (&accountVal).withdrawPointer(100)
	fmt.Println(accountVal.balance)
	fmt.Println(accountPtr.balance)

}
