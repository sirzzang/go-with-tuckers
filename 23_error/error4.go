package main

import "fmt"

type PasswordError struct {
	Len        int
	RequireLen int
}

type PasswordError3 struct {
	Len        int
	RequireLen int
}

func (err PasswordError) Error() string {
	return fmt.Sprintf(
		"Password length require %d characters or more, got %d characters", err.RequireLen, err.Len)

}

func (err PasswordError3) Error() string {
	return fmt.Sprintf(
		"Password length require %d characters or more, got %d characters", err.RequireLen, err.Len)

}

func RegisterAccount(name, password string) error {

	// TODO: 실무에서 이렇게 해도 되려나..? 한글이면?
	if len(password) < 8 {
		return PasswordError{Len: len(password), RequireLen: 8}
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
		} else {
			// TODO: 여기는 어떻게 검사해야 하지? 애초에 err가 PasswordError로 변환 안 될 가능성도 있나?
			fmt.Println("unknown error")
		}
	} else {
		fmt.Println("register success")

	}

}
