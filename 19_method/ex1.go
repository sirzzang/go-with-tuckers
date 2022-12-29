package main

import "fmt"

type Cart struct {
	productList string
}

func (c *Cart) AddProduct(product string) {
	if c.productList != "" {
		c.productList += ", "
	}
	c.productList += product
}

func (c *Cart) Clear() {
	c.productList = ""
}

func (c Cart) GetProductList() string {
	return c.productList
}

func main() {
	c := &Cart{}
	c.AddProduct("Apple")
	c.AddProduct("Kimchi")

	fmt.Println(c.GetProductList())

	c.Clear()
	c.AddProduct("Watermelon")
	fmt.Println(c.GetProductList())
}
