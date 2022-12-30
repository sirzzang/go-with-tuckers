package main

import (
	"go-with-tuckers/20_interface/fedex"
	"go-with-tuckers/20_interface/koreaPost"
)

// fedex, koreaPost에서 제공하는 API(FedexSender, KoreaPostSender)는 모두 Sender 인터페이스 구현
// 따라서 FedexSender, KoreaPostSender는 모두 Sender 타입으로 이용될 수 있음
type Sender interface {
	Send(parcel string)
}

func SendBook(book string, sender Sender) {
	sender.Send(book)
}

func main() {

	f := &fedex.FedexSender{}
	k := &koreaPost.KoreaPostSender{}
	SendBook("Tucker의 Go 언어 프로그래밍", f)
	SendBook("작별인사", k)
	SendBook("불편한 편의점", k)

}
