package main

import "fmt"

type User struct {
	Name string
	ID   string
	Age  int
}

type VIPUser struct {
	userInfo User // Embedded Type
	VIPLevel int
	Price    int
}

type VipUser struct {
	User     // Embedded Field
	VIPLevel int
	Price    int
}

func main() {

	user := User{"송이레", "sirzzang", 28}
	fmt.Println(user.Name, user.ID, user.Age)

	vipUser := VIPUser{
		User{"이레", "eraser", 28}, // Embedded Type
		3,
		250,
	}
	fmt.Println(
		vipUser.UserInfo.Name,
		vipUser.UserInfo.ID,
		vipUser.UserInfo.Age,
		vipUser.VIPLevel,
		vipUser.Price,
	)

	vipUser2 := VIPUser{
		UserInfo: user, VIPLevel: 3, Price: 250,
	}
	fmt.Println(
		vipUser2.UserInfo.Name,
		vipUser2.UserInfo.ID,
		vipUser2.UserInfo.Age,
		vipUser2.VIPLevel,
		vipUser2.Price,
	)

	vipUser3 := VipUser{
		User{"이레", "eraser", 28}, // Embedded Field
		3,
		250,
	}
	fmt.Println(
		vipUser3.Name, // 필드명 한 번에 접근 가능
		vipUser3.ID,
		vipUser3.Age,
		vipUser3.VIPLevel,
		vipUser3.Price,
	)
}
