package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
)

func main() {

	stdin := bufio.NewReader(os.Stdin)

	var num int

	answer := rand.Intn(100)
	cnt := 0

	for {
		_, err := fmt.Scanln(&num)
		cnt += 1

		if err != nil {
			fmt.Println("다시 입력하세요.", err)
			stdin.ReadString('\n') // empty buffer
		} else {
			if num > answer {
				fmt.Println("입력하신 숫자가 더 큽니다.")
			} else if num < answer {
				fmt.Println("입력하신 숫자가 더 작습니다.")
			} else {
				fmt.Printf("숫자를 맞췄습니다. 축하합니다. 시도한 횟수: %d\n", cnt)
				break
			}
		}
	}
}
