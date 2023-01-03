package main

import "fmt"

type PasswordError struct {
	Len        int
	RequireLen int
}

func (err PasswordError) Error() string {
	return "암호 길이가 짧습니다."

}

func RegisterAccount(name, password string) error {

	pwSlice := []rune(password)
	if len(pwSlice) < 8 {
		return PasswordError{Len: len(pwSlice), RequireLen: 8}
	}
	return nil
}

func main() {

	// register account
	err := RegisterAccount("sirzzang", "ㄱㄴㄷ")

	// check result
	if err != nil {
		if errType, ok := err.(PasswordError); ok {
			fmt.Printf("%v Len:%d RequrieLen: %d\n",
				errType, errType.Len, errType.RequireLen)
		}
	} else {
		fmt.Println("register success")

	}

}
