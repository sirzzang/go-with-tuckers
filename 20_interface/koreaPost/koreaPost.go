package koreaPost

import "fmt"

type KoreaPostSender struct {
}

func (k *KoreaPostSender) Send(parcel string) {
	fmt.Printf("우체국에서 %v를 보냅니다.\n", parcel)
}
