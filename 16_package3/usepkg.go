package main

import (
	"fmt"
	"go-with-tuckers/16_package3/publicpkg"
)

func main() {
	fmt.Println(publicpkg.PI)
	fmt.Println(publicpkg.Screensize)
	// fmt.Println(publicpkg.screenheight)
	publicpkg.PublicFunc()
	// publicpkg.privateFunc()

	var myint publicpkg.MyInt = 10 // TODO: 왜 MyInt라고 하면 안 되고 pkg명에서 접근해야 하지?
	fmt.Println(myint)

	mystruct := publicpkg.MyStruct{Age: 28}
	fmt.Println(mystruct)
}
