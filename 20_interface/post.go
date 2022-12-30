package main

import (
	"go-with-tuckers/20_interface/fedex"
	"go-with-tuckers/20_interface/koreaPost"
)

func SendBook(book string, sender *fedex.FedexSender) {
	sender.Send(book)
}

func main() {

	f := &fedex.FedexSender{}
	s := &koreaPost.KoreaPostSender{}
	SendBook("불편한 편의점", f)
	SendBook("The Little Prince", f)
	SendBook("크리스마스 타일", s) // panic

}
