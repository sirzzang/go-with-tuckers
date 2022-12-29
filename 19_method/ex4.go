package main

import (
	"fmt"
	"time"
	"unsafe"
)

type Courier struct {
	Name string
}

type Product struct {
	Name string
	ID   int
}

type Parcel struct {
	Pdt           *Product // pointer
	ShippedTime   time.Time
	DeliveredTime time.Time
}

// 택배회사에 속하는 메서드
func (c *Courier) SendProduct(product *Product) *Parcel {
	p := &Parcel{
		Pdt:         product,
		ShippedTime: time.Now(),
	}
	return p
}

// 소포에 속하는 메서드
func (p *Parcel) Delivered() *Product {
	p.DeliveredTime = time.Now()
	return p.Pdt
}

func main() {
	courier := &Courier{"대한통운"}
	product := &Product{"아이폰", 1}
	fmt.Println("최초 product: ", unsafe.Pointer(product))

	fmt.Println("택배 발송")
	parcel := courier.SendProduct(product) // pointer

	fmt.Println("택배 배송 완료")
	deliveredProduct := parcel.Delivered()
	fmt.Println(deliveredProduct)
	fmt.Println(deliveredProduct, unsafe.Pointer(deliveredProduct)) // 최초 product와 동일

}
