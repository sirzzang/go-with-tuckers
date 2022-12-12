package main

import "fmt"

type User struct {
	name  string
	id    string
	level int
}

type VIPUser struct {
	User
	level string // 필드명 중복
	price int
}

func main() {
	user := User{"송이레", "sirzzang", 3}
	vipUser := VIPUser{
		User{"이레이저", "eraser", 3},
		"vip",
		250,
	}

	fmt.Println(user.name, user.id, user.level)
	fmt.Println(
		vipUser.name,
		vipUser.id,
		vipUser.level, // 현재 변수 타입에 해당하는 필드 우선
		// vipUser.level,
		vipUser.User.level, // 중복된 필드의 경우 포함된 구조체명으로 접근
		vipUser.price,
	)
}
