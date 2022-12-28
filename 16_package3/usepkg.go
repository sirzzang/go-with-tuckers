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

	var myint publicpkg.MyInt = 10
	fmt.Println(myint)

	mystruct := publicpkg.MyStruct{Age: 28}
	fmt.Println(mystruct)
}
