package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
)

func UserInput() int {
	stdin := bufio.NewReader(os.Stdin)

	var num int

	_, err := fmt.Scanln(&num)
	if err != nil {
		log.Fatalf("scanner error: %v", err)
		stdin.ReadString('\n') // empty stdin buffer
	}
	return num

}

func main() {

	rand.Seed(42) // fix seed

	answer := rand.Intn(100)
	cnt := 0

	for {
		input := UserInput()
		cnt += 1
		if input > answer {
			fmt.Println("입력하신 숫자가 더 큽니다.", input, answer)
		} else if input < answer {
			fmt.Println("입력하신 숫자가 더 작습니다.", input, answer)
		} else {
			fmt.Printf("숫자를 맞췄습니다. 축하합니다. 시도한 횟수: %d\n", cnt)
			break
		}
	}
}
